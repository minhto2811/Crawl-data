package repo

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"os"
	"path/filepath"
	"strings"
	"time"

	"net/url"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type TvhlRepo interface {
	SavePractice(practice *models.TVHL, collection string) error
	Upload(title string, url1 string) (string, error)
	Clear(cutoff time.Time) error
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

    // Kiểm tra xem URL có phải là file PDF sẵn hay không
    // (Dùng strings.Contains để phòng trường hợp url có query params phía sau như .pdf?alt=media)
    isPDF := strings.Contains(strings.ToLower(url1), ".pdf")

    // 🔹 2. Tạo tên file tạm thời với extension phù hợp
    name := utils.ToSnakeCase(title)
    fileName := fmt.Sprintf("%s_%d", name, time.Now().Unix())
    tempDir := os.TempDir()
    
    ext := ".docx"
    if isPDF {
        ext = ".pdf"
    }
    
    tmpPath := filepath.Join(tempDir, fileName+ext)
    out, err := os.Create(tmpPath)
    if err != nil {
        return url1, fmt.Errorf("không tạo file tạm: %v", err)
    }
    defer out.Close()

    if _, err := io.Copy(out, resp.Body); err != nil {
        return url1, fmt.Errorf("không ghi file tạm: %v", err)
    }
    
    // Đóng file ngay sau khi ghi xong để tránh xung đột quyền đọc/ghi (lock file) ở các bước sau
    out.Close() 

    var tmpPdfPath string
    if isPDF {
        fmt.Println("🔹 Tải trực tiếp PDF thành công (Không cần convert)")
        tmpPdfPath = tmpPath
    } else {
        fmt.Println("🔹 Tải docx thành công, bắt đầu convert...")
        tmpPdfPath, err = utils.ConvertDocxToPDF(tmpPath, tempDir, 0)
        if err != nil {
            return url1, fmt.Errorf("không chuyển đổi được file pdf: %v", err)
        }
        fmt.Println("🔹 Convert sang PDF thành công")
    }

    fmt.Println("🔹 Upload lên server")
    // 🔹 3. Upload lên Firebase Storage
    objectPath := fmt.Sprintf("TVHL/%s", fileName+".pdf")
    wc := r.storage.Bucket(r.bucketName).Object(objectPath).NewWriter(r.ctx)
    uuid := uuid.New().String()
    wc.ContentType = "application/pdf"
    wc.Metadata = map[string]string{
        "firebaseStorageDownloadTokens": uuid,
    }

    // Mở file tạm PDF để upload
    tmpFile, err := os.Open(tmpPdfPath)
    if err != nil {
        return url1, fmt.Errorf("không mở file tạm: %v", err)
    }
    defer tmpFile.Close()

    if _, err := io.Copy(wc, tmpFile); err != nil {
        wc.Close()
        return url1, fmt.Errorf("lỗi upload : %v", err)
    }

    if err := wc.Close(); err != nil {
        return url1, fmt.Errorf("lỗi đóng writer: %v", err)
    }

    // Tự động dọn dẹp file tạm sau khi chạy xong để đỡ rác server
    defer os.Remove(tmpPath)
    if tmpPdfPath != tmpPath {
        defer os.Remove(tmpPdfPath)
    }

    // 🔹 4. Tạo URL công khai
    escapedPath := url.PathEscape(objectPath)
    newURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", r.bucketName, escapedPath, uuid)

    fmt.Println("🔹 Upload thành công")
    return newURL, nil
}

func (r *TvhlRepoImp) Clear(cutoff time.Time) error {
	collections := []string{"literatures", "naturalScience", "english", "physics", "biology", "history", "chemistry", "geography", "civics"}
	for _, collection := range collections {
		snapshot, err := r.client.Collection(collection).Where("lastModified", "<=", cutoff).Documents(r.ctx).GetAll()
		if err != nil {
			return err
		}
		fmt.Printf("Tài liệu %s trước tháng %d năm %d: %d\n", collection, cutoff.Month(), cutoff.Year(), len(snapshot))
		for _, doc := range snapshot {
			fileURL, ok := doc.Data()["url"].(string)
			if !ok || fileURL == "" {
				return fmt.Errorf("document %s không có field url hợp lệ", doc.Ref.ID)
			}
			start := strings.Index(fileURL, "/o/")
			if start == -1 {
				return fmt.Errorf("URL không hợp lệ: %s", fileURL)
			}
			start += len("/o/")
			end := strings.Index(fileURL, "?")
			if end == -1 {
				end = len(fileURL)
			}
			path := fileURL[start:end]
			path = strings.ReplaceAll(path, "%2F", "/")
			obj := r.storage.Bucket(r.bucketName).Object(path)
			if err := obj.Delete(r.ctx); err != nil {
				if errors.Is(err, storage.ErrObjectNotExist) {
					fmt.Println("⚠️ File không tồn tại")
				} else {
					return fmt.Errorf("lỗi khi xóa file %q: %v", path, err)
				}

			}
			fmt.Printf("🗑️ Đã xóa file: %s\n", path)
			fmt.Printf("Removing document %s\n", doc.Ref.ID)
			_, err := r.client.Collection(collection).Doc(doc.Ref.ID).Delete(r.ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
