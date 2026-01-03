package crawl

import (
	"fmt"
	"log"
	"mxgk/crawl/models"
	"mxgk/crawl/repo"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

const (
	minTime = "02/01/2026"
	maxTime = "03/01/2026"
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
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "vi-VN,vi;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func getDocumentWithRod(url string) (*goquery.Document, error) {
    l := launcher.New().
        Headless(true).
        UserDataDir("rod_data")

    controlURL, err := l.Launch()
    if err != nil {
        return nil, err
    }

    browser := rod.New().ControlURL(controlURL).MustConnect()
    defer browser.MustClose()

    // SỬA TẠI ĐÂY: Truyền 'browser' (kiểu *rod.Browser) 
    // Hàm này sẽ trả về một '*rod.Page' đã được cài đặt stealth
    page := stealth.MustPage(browser)

    // Thực hiện truy cập
    err = page.Navigate(url)
    if err != nil {
        return nil, err
    }

    // Đợi trang web vượt qua cơ chế check bot (thường mất vài giây)
    fmt.Println("Đang đợi xác thực trình duyệt...")
    page.MustWaitIdle()
    
    // Đợi thêm một chút nếu cần (trang bạn gửi có setTimeout 5s)
    time.Sleep(6 * time.Second) 

    html, err := page.HTML()
    if err != nil {
        return nil, err
    }

    return goquery.NewDocumentFromReader(strings.NewReader(html))
}
