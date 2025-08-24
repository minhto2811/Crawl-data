package main

import (
	"context"
	"log"
	"mxgk/crawl/crawl"
	"mxgk/crawl/repo"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)



var practiceRepo repo.PracticeRepo

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	client, err := firestore.NewClient(ctx, "dehay-73822", sa)
	if err != nil {
		log.Fatalf("Lỗi tạo Firestore client: %v", err)
	}
	defer client.Close()

	practiceRepo = repo.NewPracticeRepo(client, ctx)
	crawl.SetRepo(practiceRepo)
	//crawl.CrawlLiterature()
	crawl.CrawlMath()
}

