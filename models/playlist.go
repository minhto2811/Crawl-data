package models

import (
	"time"
)


type Playlist struct{
	Id		 string `firestore:"id"`
	Title	 string `firestore:"title"`
	Thumbnail string `firestore:"thumbnail"`
	Grade	 string `firestore:"grade"`
	LastModified time.Time  `firestore:"last_modified"`
	Count	 int    `firestore:"count"`
}