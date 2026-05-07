package repo

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"mxgk/crawl/models"
	"time"
)

type VideoRepo interface {
	UploadVideos(videos []models.Video) error
	UploadPlaylist(playlist models.Playlist) error
	GetLastModifiedAndCountPlaylist(playlistID string) (time.Time, error)
}

type VideoRepoImp struct {
	client *firestore.Client
	ctx    context.Context
}

func NewVideoRepo(client *firestore.Client, ctx context.Context) VideoRepo {
	return &VideoRepoImp{client: client, ctx: ctx}
}

func (r *VideoRepoImp) UploadVideos(videos []models.Video) error {
	for _, video := range videos {
		_, err := r.client.Collection("videos").Doc(video.Id).Set(r.ctx, video)
		if err != nil {
			fmt.Printf("Lỗi thêm video: %s, lỗi: %v\n", video.Title, err)
			continue
		}
		fmt.Printf("Đã thêm video: %s\n", video.LastModified)
	}
	return nil
}

func (r *VideoRepoImp) UploadPlaylist(playlists models.Playlist) error {
	query := r.client.Collection("videos").Where("playlist", "==", playlists.Id)
	aggregationQuery := query.NewAggregationQuery().WithCount("all_videos")

	// 2. Thực thi query
	results, err := aggregationQuery.Get(r.ctx)
	if err != nil {
		return err
	}

	// 3. Trích xuất giá trị từ map kết quả
	countValue, ok := results["all_videos"]
	if !ok {
		return fmt.Errorf("không tìm thấy kết quả đếm")
	}

	// 4. Ép kiểu về int64 (Firestore trả về kiểu này)
	count, ok := countValue.(int64)
	if !ok {
		return fmt.Errorf("count không phải int64")
	}
	playlists.Count = int(count)
	_, err1 := r.client.Collection("playlists").Doc(playlists.Id).Set(r.ctx, playlists)
	if err1 != nil {
		fmt.Printf("Lỗi thêm playlist: %s, lỗi: %v\n", playlists.Title, err1)
		return err1
	}
	fmt.Println("Đã thêm playlist: ", playlists.LastModified)
	return nil
}

func (r *VideoRepoImp) GetLastModifiedAndCountPlaylist(playlistID string) (time.Time, error) {
	iter := r.client.Collection("videos").
		Where("playlist", "==", playlistID).
		OrderBy("lastModified", firestore.Desc).
		Limit(1).
		Documents(r.ctx) // Use .Documents() for queries

	defer iter.Stop()

	// 2. Get the first result
	doc, err := iter.Next()
	if err != nil {
		return time.Time{}, err
	}
	return doc.Data()["lastModified"].(time.Time), nil
}
