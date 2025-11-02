package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/firestore"
)

// Structs -------------------------

type Question struct {
	Question    string   `json:"question" firestore:"question,omitempty"`
	Answers     []string `json:"answers" firestore:"answers,omitempty"`
	IndexAnswer int      `json:"indexAnswer" firestore:"indexAnswer"`
	Solution    string   `json:"solution" firestore:"solution,omitempty"`
}

type Test struct {
	ID          string     `json:"id" firestore:"id,omitempty"`
	Name        string     `json:"name" firestore:"name,omitempty"`
	Description string     `json:"description" firestore:"description,omitempty"`
	Grade       string     `json:"grade" firestore:"grade,omitempty"`
	Questions   []Question `json:"questions" firestore:"questions,omitempty"`
}


// Interface ----------------------

type TestRepo interface {
	Update() error
}

// Implementation -----------------

type TestRepoImp struct {
	client *firestore.Client
	ctx    context.Context
}

// Constructor --------------------

func NewTestRepo(client *firestore.Client, ctx context.Context) TestRepo {
	return &TestRepoImp{client: client, ctx: ctx}
}

// Update method ------------------

func (r *TestRepoImp) Update() error {
	// Đọc file test.json
	data, err := os.ReadFile("D:/Golang_Projects/Crawl-data/test.json")
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	// Parse JSON -> struct
	var test Test
	if err := json.Unmarshal(data, &test); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Tạo document test
	testRef := r.client.Collection("tests").NewDoc()
	_, err = testRef.Set(r.ctx, map[string]interface{}{
		"id":          testRef.ID,
		"name":        test.Name,
		"description": test.Description,
		"grade":       test.Grade,
		"lastModified": time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to upload test: %v", err)
	}

	// Tạo subcollection questions
	for i, q := range test.Questions {
		docID := fmt.Sprintf("q%02d", i+1)
		_, err := testRef.Collection("questions").Doc(docID).Set(r.ctx, q)
		if err != nil {
			return fmt.Errorf("failed to upload question %d: %v", i+1, err)
		}
	}

	fmt.Printf("✅ Uploaded test '%s' with %d questions\n", test.Name, len(test.Questions))
	return nil
}
