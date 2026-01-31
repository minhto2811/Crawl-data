package repo

import (
	"context"
	"fmt"
	"io"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"os"
	"path/filepath"
	"time"

	"net/url"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type TvhlRepo interface {
	SavePractice(practice *models.TVHL, collection string) error
	Upload(title string, url1 string) (string, error)
}

type TvhlRepoImp struct {
	client     *firestore.Client
	ctx        context.Context
	storage    *storage.Client
	bucketName string
}

func NewTvhlRepo(client *firestore.Client, ctx context.Context, storage *storage.Client, bucketName string) TvhlRepo {
	return &TvhlRepoImp{client: client, ctx: ctx, storage: storage, bucketName: bucketName}
}

func (r *TvhlRepoImp) SavePractice(practice *models.TVHL, collection string) error {
	_, err := r.client.Collection(collection).NewDoc().Create(r.ctx, practice)
	return err
}


func (r *TvhlRepoImp) Upload(title string, url1 string) (string, error) {
	fmt.Println("🔹 Bắt đầu tải tài liệu")
	resp, err := utils.CreateRequest(url1, "GET")
	if err != nil {
		return url1, fmt.Errorf("không tải được %s: %v", url1, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return url1, fmt.Errorf("HTTP lỗi: %d", resp.StatusCode)
	}

	// 🔹 2. Tạo tên file tạm thời
	name := utils.ToSnakeCase(title)
	fileName := fmt.Sprintf("%s_%d", name, time.Now().Unix())
	tempDir := os.TempDir()
	tmpPath := filepath.Join(tempDir, fileName+".docx")
	out, err := os.Create(tmpPath)
	if err != nil {
		return url1, fmt.Errorf("không tạo file tạm: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return url1, fmt.Errorf("không ghi file tạm: %v", err)
	}
	fmt.Println("🔹 Tải docx thành công")

	tmpPdfPath, err1 := utils.ConvertDocxToPDF(tmpPath, tempDir, 0)

	if err1 != nil {
		return url1, fmt.Errorf("không chuyển đổi được file pdf: %v", err1)
	}
	fmt.Println("🔹 Convert sang PDF thành công")
	fmt.Println("🔹 Upload lên server")
	// 🔹 3. Upload lên Firebase Storage
	objectPath := fmt.Sprintf("TVHL/%s", fileName+".pdf")
	wc := r.storage.Bucket(r.bucketName).Object(objectPath).NewWriter(r.ctx)
	defer wc.Close()
	uuid := uuid.New().String()
	wc.ContentType = "application/pdf"
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid,
	}

	// Mở lại file tạm để upload
	tmpFile, err := os.Open(tmpPdfPath)
	if err != nil {
		return url1, fmt.Errorf("không mở file tạm: %v", err)
	}
	defer tmpFile.Close()

	if _, err := io.Copy(wc, tmpFile); err != nil {
		return url1, fmt.Errorf("lỗi upload : %v", err)
	}

	if err := wc.Close(); err != nil {
		return url1, fmt.Errorf("lỗi đóng writer: %v", err)
	}

	// 🔹 4. Tạo URL công khai
	escapedPath := url.PathEscape(objectPath)
	newURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", r.bucketName, escapedPath, uuid)

	fmt.Println("🔹 Upload thành công")
	return newURL, nil
}
