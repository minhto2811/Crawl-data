package main

import (
	"context"
	"log"
	"mxgk/crawl/crawl"
	"mxgk/crawl/repo"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
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

	stor1, err := storage.NewClient(ctx, sa)
	if err != nil {
		// TODO: Handle error.
	}

	practiceRepo = repo.NewPracticeRepo(client, ctx, stor1, "dehay-73822.firebasestorage.app")
	crawl.SetRepo(practiceRepo)

	// crawl.CrawlVideo("PLXYp7Odn5ED8jj_ROzHTVt5H4NNYZtY66","up")

	crawl.CrawlMath()
}
