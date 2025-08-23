package repo

import (
	"context"
	"fmt"
	"mxgk/crawl/models"

	"cloud.google.com/go/firestore"
)

type PracticeRepo interface {
	SavePractice(practice *models.Practice,collection string) error
	Update() error
	Remove() error
}

type PracticeRepoImp struct {
	client *firestore.Client
	ctx    context.Context
}

func NewPracticeRepo(client *firestore.Client, ctx context.Context) PracticeRepo {
	return &PracticeRepoImp{client: client, ctx: ctx}
}

func (r *PracticeRepoImp) SavePractice(practice *models.Practice, collection string) error {
	_, err := r.client.Collection(collection).NewDoc().Create(r.ctx, practice)
	return err
}

func (r *PracticeRepoImp) Update() error {
	snapshot, err := r.client.Collection("practices").Documents(r.ctx).GetAll()
	if err != nil {
		return err
	}
	for _, doc := range snapshot {
		if doc.Data()["subject"] != nil {
			continue
		}
		_, err := r.client.Collection("practices").Doc(doc.Ref.ID).Set(r.ctx, map[string]interface{}{
			"subject": "math",
		}, firestore.MergeAll)
		if err != nil {
			fmt.Printf("Error: %s\n", doc.Ref.ID)
			return err
		}
		fmt.Printf("Updated document %s\n", doc.Ref.ID)
	}

	return nil
}

func (r *PracticeRepoImp) Remove() error {
	snapshot, err := r.client.Collection("practices").Where("type", "==", "sgk").Documents(r.ctx).GetAll()
	if err != nil {
		return err
	}
	for _, doc := range snapshot {
		fmt.Printf("Removing document %s\n", doc.Ref.ID)
		_, err := r.client.Collection("practices").Doc(doc.Ref.ID).Delete(r.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}


