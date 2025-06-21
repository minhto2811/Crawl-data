package repo

import (
	"context"
	"fmt"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"strings"

	"cloud.google.com/go/firestore"
)

type PracticeRepo interface {
	SavePractice(practice *models.Practice) error
	Update() error
}

type PracticeRepoImp struct {
	client *firestore.Client
	ctx    context.Context
}

func NewPracticeRepo(client *firestore.Client, ctx context.Context) PracticeRepo {
	return &PracticeRepoImp{client: client, ctx: ctx}
}

func (r *PracticeRepoImp) SavePractice(practice *models.Practice) error {
	_, err := r.client.Collection("practices").NewDoc().Create(r.ctx, practice)
	return err
}

func (r *PracticeRepoImp) Update() error {
	snapshot, err := r.client.Collection("practices").Documents(r.ctx).GetAll()
	if err != nil {
		return err
	}
	for _, doc := range snapshot {
		if _, ok := doc.Data()["keyWords"]; !ok {
			var practice models.Practice
			if err := doc.DataTo(&practice); err != nil {
				return err
			}
			withoutDiacritics := utils.RemoveDiacritics(practice.Title)
			lowered := strings.ToLower(withoutDiacritics)
			keyWords := strings.Fields(lowered)
			fmt.Printf("Updating document %s with keyWords: %v\n", doc.Ref.ID, keyWords)
			_, err := r.client.Collection("practices").Doc(doc.Ref.ID).Set(r.ctx, map[string]interface{}{
				"keyWords": keyWords,
			}, firestore.MergeAll)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


