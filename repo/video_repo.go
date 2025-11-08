package repo

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxgk/crawl/models"
	"time"
)

type VideoRepo interface {
	UploadVideos(videos []models.Video) error
	UploadPlaylist(playlist models.Playlist) error
	GetLastModifiedPlaylist(playlistID string) (time.Time, error)
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
	_, err := r.client.Collection("playlists").Doc(playlists.Id).Set(r.ctx, playlists)
	if err != nil {
		fmt.Printf("Lỗi thêm playlist: %s, lỗi: %v\n", playlists.Title, err)
		return err
	}
	fmt.Printf("Đã thêm playlist: %s\n", playlists.Title)
	return nil
}

func (r *VideoRepoImp) GetLastModifiedPlaylist(playlistID string) (time.Time, error) {
	doc, err := r.client.Collection("playlists").Doc(playlistID).Get(r.ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			fmt.Println("Document không tồn tại:", playlistID)
			return time.Time{}, nil // trả về default time
		}
		return time.Time{}, err
	}
	var pl models.Playlist
	err = doc.DataTo(&pl)
	if err != nil {
		return time.Time{}, err
	}
	return pl.LastModified, nil
}
