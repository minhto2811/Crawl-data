package repo

import (
	"context"
	"fmt"
	"io"
	"mxgk/crawl/models"
	"strings"
	"time"

	"net/http"
	"path/filepath"

	"net/url"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type PracticeRepo interface {
	SavePractice(practice *models.Practice, collection string) error
	Update() error
	Remove() error
	Backup() error
	Tutorial(videos []models.Video) error
}

type PracticeRepoImp struct {
	client     *firestore.Client
	ctx        context.Context
	storage    *storage.Client
	bucketName string
}

func NewPracticeRepo(client *firestore.Client, ctx context.Context, storage *storage.Client, bucketName string) PracticeRepo {
	return &PracticeRepoImp{client: client, ctx: ctx, storage: storage, bucketName: bucketName}
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

func (r *PracticeRepoImp) Backup() error {
	snapshot, err := r.client.Collection("practices").Documents(r.ctx).GetAll()
	if err != nil {
		return err
	}
	fmt.Printf("Đã tìm được %d bản ghi\n", len(snapshot))
	for index, doc := range snapshot {
		fmt.Printf("Bản ghi thứ %d/%d\n", index+1, len(snapshot))
		var p models.Practice
		if err := doc.DataTo(&p); err != nil {
			fmt.Printf("Lỗi đọc dữ liệu: %v\n", err)
			continue
		}

		if p.Url == "" {
			continue
		}

		if strings.Contains(p.Url, "firebasestorage") {
			fmt.Printf("Đã backup rồi, bỏ qua: %s\n", p.Url)
			continue
		}

		// 🔹 1. Tải file từ URL cũ
		resp, err := http.Get(p.Url)
		if err != nil {
			fmt.Printf("Không tải được %s: %v\n", p.Url, err)
			continue
		}
		defer resp.Body.Close()

		// 🔹 2. Tạo tên file trong Firebase Storage
		fileName := filepath.Base(p.Url)
		if fileName == "" {
			fileName = fmt.Sprintf("%s_%d", doc.Ref.ID, time.Now().Unix())
		}
		objectPath := fmt.Sprintf("math/%s", fileName)

		// 🔹 3. Upload file lên Storage
		wc := r.storage.Bucket(r.bucketName).Object(objectPath).NewWriter(r.ctx)
		uuid := uuid.New().String()
		wc.Metadata = map[string]string{
			"firebaseStorageDownloadTokens": uuid, // <--- HERE
		}
		if _, err := io.Copy(wc, resp.Body); err != nil {
			fmt.Printf("Lỗi upload %s: %v\n", fileName, err)
			wc.Close()
			continue
		}
		if err := wc.Close(); err != nil {
			fmt.Printf("Lỗi đóng writer: %v\n", err)
			continue
		}

		// 🔹 4. Lấy URL công khai (hoặc signed URL nếu bạn muốn bảo mật)
		escapedPath := url.PathEscape(objectPath)
		newURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", r.bucketName, escapedPath, uuid)

		// 🔹 5. Cập nhật lại document Firestore
		_, err = doc.Ref.Update(r.ctx, []firestore.Update{
			{Path: "url", Value: newURL},
		})
		if err != nil {
			fmt.Printf("Lỗi cập nhật Firestore: %v\n", err)
			continue
		}

		fmt.Printf("✅ Đã backup file: %s\n", newURL)
	}
	return nil
}

func (r *PracticeRepoImp) Tutorial(videos []models.Video) error {
	for _, video := range videos {
		_, err := r.client.Collection("tutorials").NewDoc().Create(r.ctx, video)
		if err != nil {
			fmt.Printf("Lỗi thêm video: %s, lỗi: %v\n", video.Title, err)
			continue
		}
		fmt.Printf("Đã thêm video: %s\n", video.Title)
	}
	return nil
}
