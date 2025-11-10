package repo

import (
	"context"
	"fmt"
	"io"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"os"
	"time"

	"net/http"
	"path/filepath"

	"net/url"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type HocMaiRepo interface {
	SavePractice(practice *models.HocMai, collection string) error
	Upload(title string, url1 string, cookie string) (string, error)
}

type HocMaiRepoImp struct {
	client     *firestore.Client
	ctx        context.Context
	storage    *storage.Client
	bucketName string
}

func NewHocMaiRepo(client *firestore.Client, ctx context.Context, storage *storage.Client, bucketName string) HocMaiRepo {
	return &HocMaiRepoImp{client: client, ctx: ctx, storage: storage, bucketName: bucketName}
}

func (r *HocMaiRepoImp) SavePractice(practice *models.HocMai, collection string) error {
	_, err := r.client.Collection(collection).NewDoc().Create(r.ctx, practice)
	return err
}

func (r *HocMaiRepoImp) Upload(title string, url1 string, cookie string) (string, error) {
	// 🔹 1. Tải file từ URL cũ
	req, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		return url1, fmt.Errorf("tạo request lỗi: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/142.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "vi")
	req.Header.Set("Cookie", cookie)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return url1, fmt.Errorf("không tải được %s: %v", url1, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return url1, fmt.Errorf("HTTP lỗi: %d", resp.StatusCode)
	}

	// 🔹 2. Tạo tên file tạm thời
	name := utils.ToSnakeCase(title)
	fileName := fmt.Sprintf("%s_%d.pdf", name, time.Now().Unix())

	tmpPath := filepath.Join(os.TempDir(), fileName)
	out, err := os.Create(tmpPath)
	if err != nil {
		return url1, fmt.Errorf("không tạo file tạm: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return url1, fmt.Errorf("không ghi file tạm: %v", err)
	}

	// 🔹 3. Upload lên Firebase Storage
	objectPath := fmt.Sprintf("literature/%s", fileName)
	wc := r.storage.Bucket(r.bucketName).Object(objectPath).NewWriter(r.ctx)
	uuid := uuid.New().String()
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid,
	}

	// Mở lại file tạm để upload
	tmpFile, err := os.Open(tmpPath)
	if err != nil {
		wc.Close()
		return url1, fmt.Errorf("không mở file tạm: %v", err)
	}
	defer tmpFile.Close()

	if _, err := io.Copy(wc, tmpFile); err != nil {
		wc.Close()
		return url1, fmt.Errorf("lỗi upload %s: %v", fileName, err)
	}

	if err := wc.Close(); err != nil {
		return url1, fmt.Errorf("lỗi đóng writer: %v", err)
	}

	// 🔹 4. Tạo URL công khai
	escapedPath := url.PathEscape(objectPath)
	newURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", r.bucketName, escapedPath, uuid)

	fmt.Println("✅ Upload thành công:", newURL)
	return newURL, nil
}
