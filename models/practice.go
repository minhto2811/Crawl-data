package models

import "time"

type Practice struct {
	Title        string   `firestore:"title"`
	Grade        string   `firestore:"grade"`
	Subject      string   `firestore:"subject"`
	Url          string   `firestore:"url"`
	LastModified time.Time `firestore:"lastModified"`
	KeyWords     []string `firestore:"keyWords,omitempty"` // Từ khóa tìm kiếm, có thể để trống
}