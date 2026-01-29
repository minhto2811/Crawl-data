package repo

import (
	"context"
	"fmt"
	"mxgk/crawl/models"
	"time"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VideoRepo interface {
	UploadVideos(videos []models.Video) error
	UploadPlaylist(playlist models.Playlist) error
	GetLastModifiedAndCountPlaylist(playlistID string) (time.Time, int, error)
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
		fmt.Printf("Đã thêm video: %s\n", video.Title)
	}
	return nil
}

func (r *VideoRepoImp) UploadPlaylist(playlists models.Playlist) error {
	query := r.client.
		Collection("videos").
		Where("playlist", "==", playlists.Id)

	aggregationQuery := query.NewAggregationQuery().WithCount("all")
	results, err := aggregationQuery.Get(r.ctx)
	if err != nil {
		panic(err)
	}
	count, ok := results["all"]
	if ok {
		countValue := count.(*firestorepb.Value)
		playlists.Count = int(countValue.GetIntegerValue())
	}

	_, err1 := r.client.Collection("playlists").Doc(playlists.Id).Set(r.ctx, playlists)
	if err1 != nil {
		fmt.Printf("Lỗi thêm playlist: %s, lỗi: %v\n", playlists.Title, err1)
		return err1
	}
	fmt.Printf("Đã thêm playlist: %s\n", playlists.Title)
	return nil
}

func (r *VideoRepoImp) GetLastModifiedAndCountPlaylist(playlistID string) (time.Time, int, error) {
	doc, err := r.client.Collection("playlists").Doc(playlistID).Get(r.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			fmt.Println("Document không tồn tại:", playlistID)
			return time.Time{}, 0, nil // trả về default time
		}
		return time.Time{}, 0, err
	}
	var pl models.Playlist
	err = doc.DataTo(&pl)
	if err != nil {
		return time.Time{}, 0, err
	}
	return pl.LastModified, pl.Count, nil
}
