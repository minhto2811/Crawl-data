package models

import (
	"time"
)

type Video struct {
	Id		   	 string    `firestore:"id"`
	Title        string    `firestore:"title"`
	URL          string    `firestore:"url"`
	Grade        string    `firestore:"grade"`
	LastModified time.Time `firestore:"lastModified"`
	Playlist     string    `firestore:"playlist"`
}
