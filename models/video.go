package models

import (
	"time"
)

type Video struct {
	Title        string    `firestore:"title"`
	URL          string    `firestore:"url"`
	Grade        string    `firestore:"grade"`
	LastModified time.Time `firestore:"lastModified"`
	Playlist     string    `firestore:"playlist"`
}
