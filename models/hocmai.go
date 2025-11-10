package models

import "time"

type HocMai struct {
	Title        string    `firestore:"title"`
	Grade        string    `firestore:"grade"`
	Url          string    `firestore:"url"`
	LastModified time.Time `firestore:"lastModified"`
}
