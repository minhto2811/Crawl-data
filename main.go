package main

import (
	"context"
	"fmt"
	"log"
	"mxgk/crawl/crawl"
	"mxgk/crawl/repo"
	// "time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

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
		fmt.Println("Lỗi tạo Storage client:", err)
		return
		// TODO: Handle error.
	}

	practiceRepo := repo.NewPracticeRepo(client, ctx, stor1, "dehay-73822.firebasestorage.app")
	videoRepo := repo.NewVideoRepo(client, ctx)
	tvhlRepo := repo.NewTvhlRepo(client, ctx, stor1, "dehay-73822.firebasestorage.app")
	crawl.SetRepo(practiceRepo, videoRepo, tvhlRepo)

	// crawl.CrawlVideo()
	//crawl.CrawlMath()
	crawl.CrawlTVHL()

	// cutoff := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	// err1 :=	crawl.ClearMath(cutoff)
	// if err1 != nil {
	// 	fmt.Println("Lỗi xóa tài liệu toán:", err1)
	// 	return
	// }
	// err2 := crawl.ClearTVHL(cutoff)
	// if err2 != nil {
	// 	fmt.Println("Lỗi xóa tài liệu TVHL:", err2)
	// 	return
	// }
}
