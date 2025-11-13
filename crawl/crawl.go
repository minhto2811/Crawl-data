package crawl

import (
	"fmt"
	"log"
	"mxgk/crawl/models"
	"mxgk/crawl/repo"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	minTime = "12/11/2025"
	maxTime = "13/11/2025"
)

var grade = "g8"
var g = 8
var collection = "practices"

var pRep repo.PracticeRepo
var vRep repo.VideoRepo
var hmRep repo.HocMaiRepo

func SetRepo(practiceRepo repo.PracticeRepo, videoRepo repo.VideoRepo, hocmaiRepo repo.HocMaiRepo) {
	pRep = practiceRepo
	vRep = videoRepo
	hmRep = hocmaiRepo
}




func BackUp() error {
	return pRep.Backup()
}

func updatePracticeToFirestore(practice models.Practice) error {
	log.Printf("Url cũ: %s", practice.Url)
	practice.Url = pRep.Upload(practice.Url)
	log.Printf("Url mới: %s", practice.Url)

	// Cập nhật lại document Firestore
	err := pRep.SavePractice(&practice, collection)
	if err != nil {
		log.Printf("Lỗi cập nhật practice: %v", err)
		return err
	}

	log.Printf("✅ Đã cập nhật practice: %s", practice.Title)
	return nil
}

func getDocument(url string) (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return goquery.NewDocumentFromReader(res.Body)
}
