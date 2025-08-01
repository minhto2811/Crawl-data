package main

import (
	"context"
	"fmt"
	"log"
	"mxgk/crawl/models"
	"mxgk/crawl/repo"
	"mxgk/crawl/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/PuerkitoBio/goquery"
	"google.golang.org/api/option"
)

const (
	minTime = "31/07/2025"
	maxTime = "1/08/2025" //Nếu bằng ngày hôm nay tức là đã crawl
)

var grade = "g8"
var g = 8

var practiceRepo repo.PracticeRepo

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	client, err := firestore.NewClient(ctx, "dehay-73822", sa)
	if err != nil {
		log.Fatalf("Lỗi tạo Firestore client: %v", err)
	}
	defer client.Close()

	practiceRepo = repo.NewPracticeRepo(client, ctx)
	CrawlUp()
	CrawlHsp()
	CrawlGrade()
}

func CrawlUp() {
	grade = "up"
	var sources = []models.Input{
		{
			Url:  "https://toanmath.com/de-thi-thu-thpt-mon-toan/page/",
			Type: "mock",
		},
		{
			Url:  "https://toanmath.com/de-thi-thpt-mon-toan-chinh-thuc/page/",
			Type: "official",
		},
		{
			Url:  "https://toanmath.com/de-danh-gia-nang-luc-mon-toan/page/",
			Type: "aptitude",
		},
		{
			Url:  "https://toanmath.com/tai-lieu-on-thi-thpt-mon-toan/page/",
			Type: "materials",
		},
	}
	for _, source := range sources {
		AutoCrawl(source)
	}
}

func CrawlHsp() {
	grade = "hsp"
	var sources = []models.Input{
		{
			Url:  "https://thcs.toanmath.com/de-thi-tuyen-sinh-lop-10-mon-toan/page/",
			Type: "exam",
		},
		{
			Url:  "https://thcs.toanmath.com/tai-lieu-toan-on-thi-vao-lop-10/page/",
			Type: "materials",
		},
	}
	for _, source := range sources {
		AutoCrawl(source)
	}
}

func CrawlGrade() {
	for i := 6; i <= 12; i++ {
		grade = fmt.Sprintf("g%d", i)
		g = i
		var sources = []models.Input{}
		if i <= 9 {
			var s1 = []models.Input{
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/tai-lieu-toan-%d/page/", g),
					Type: "materials",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/de-cuong-on-tap-toan-%d/page/", g),
					Type: "outline",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/de-thi-giua-hk1-toan-%d/page/", g),
					Type: "midterm1",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/de-thi-hk1-toan-%d/page/", g),
					Type: "final1",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/de-thi-giua-hk2-toan-%d/page/", g),
					Type: "midterm2",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/de-thi-hk2-toan-%d/page/", g),
					Type: "final2",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/khao-sat-chat-luong-toan-%d/page/", g),
					Type: "assessment",
				},
				{
					Url:  fmt.Sprintf("https://thcs.toanmath.com/de-thi-hsg-toan-%d/page/", g),
					Type: "gifted",
				},
			}
			sources = append(sources, s1...)
		} else {
			var s2 = []models.Input{
				{
					Url:  fmt.Sprintf("https://toanmath.com/de-cuong-on-tap-toan-%d/page/", g),
					Type: "outline",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/de-thi-giua-hk1-toan-%d/page/", g),
					Type: "midterm1",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/de-thi-hk1-toan-%d/page/", g),
					Type: "final1",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/de-thi-giua-hk2-toan-%d/page/", g),
					Type: "midterm2",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/de-thi-hk2-toan-%d/page/", g),
					Type: "final2",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/khao-sat-chat-luong-toan-%d/page/", g),
					Type: "assessment",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/de-thi-hsg-toan-%d/page/", g),
					Type: "gifted",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/giao-an-toan-%d/page/", g),
					Type: "lessonPlan",
				},
				{
					Url:  fmt.Sprintf("https://toanmath.com/tips-giai-toan-%d/page/", g),
					Type: "tips",
				},
			}
			sources = append(sources, s2...)
		}

		for _, source := range sources {
			AutoCrawl(source)
		}
	}
}

func AutoCrawl(source models.Input) {
	var listPractice []models.Practice
	pageCount := 1
	for i := 1; i <= pageCount; i++ {
		url := source.Url + fmt.Sprint(i)
		fmt.Println("Đang crawl: " + url)
		fmt.Printf("Đang lấy trang %d/%d\n", i, pageCount)
		p, newPageCount, err := getListPractice(url, source.Type)
		if err != nil {
			fmt.Println("err: " + err.Error())
			break
		}
		listPractice = append(listPractice, p...)
		if newPageCount > pageCount {
			pageCount = newPageCount
		}
	}
	for _, practice := range listPractice {
		fmt.Println("Title: " + practice.Title)
		fmt.Println("Grade: " + practice.Grade)
		fmt.Println("Type: " + practice.Type)
		fmt.Println("Url: " + practice.Url)
		fmt.Println("Last Modified: " + practice.LastModified.Format("02/01/2006"))
		UpdatePracticeToFirestore(practice)
		fmt.Println("----------------------------")
	}
}

func UpdatePracticeToFirestore(practice models.Practice) error {
	err := practiceRepo.SavePractice(&practice)
	if err != nil {
		log.Printf("Lỗi cập nhật practice: %v", err)
		return err
	}

	log.Printf("✅ Đã cập nhật practice: %s", practice.Title)
	return nil
}

func getListPractice(url string, type1 string) ([]models.Practice, int, error) {
	doc, err := getDocument(url)
	if err != nil {
		return nil, 1, err
	}

	var practices []models.Practice
	min, _ := convertToTimestamp(minTime)
	max, _ := convertToTimestamp(maxTime)
	maxPage := 1
	doc.Find(".page-numbers").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		num, err := strconv.Atoi(text)
		if err == nil && num > maxPage {
			maxPage = num
		}
	})
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		a := s.Find("h3.entry-title a")
		title := strings.TrimSpace(a.Text())
		link := strings.TrimSpace(a.AttrOr("href", ""))
		b := s.Find("div.mh-meta span.entry-meta-date a")
		dateStr := strings.TrimSpace(b.Text())

		time, err := convertToTimestamp(dateStr)
		if err != nil {
			fmt.Println("❌ Lỗi convert timestamp:", err)
			return
		}
		if time.Unix() < min.Unix() || time.Unix() > max.Unix() {
			fmt.Println("⚠️ Bỏ qua practice:", title)
			return // bỏ qua nếu thời gian nhỏ hơn minTime
		}

		fmt.Println("----------------------------")

		doc1, err1 := getDocument(link)
		if err1 != nil {
			fmt.Println("❌ Error fetching detail page:", err1)
			return
		}

		// Tìm link PDF đầu tiên
		pdfLink := ""
		doc1.Find("a[href$='.pdf']").EachWithBreak(func(i1 int, s1 *goquery.Selection) bool {
			pdfLink = strings.TrimSpace(s1.AttrOr("href", ""))
			return false // dừng sau khi lấy được link đầu tiên
		})

		if pdfLink == "" {
			fmt.Println("⚠️ Không tìm thấy PDF:", title)
			return
		}

		withoutDiacritics := utils.RemoveDiacritics(title)
		lowered := strings.ToLower(withoutDiacritics)
		keyWords := strings.Fields(lowered)

		practice := models.Practice{
			Title:        title,
			Grade:        grade,
			Type:         type1,
			Url:          pdfLink,
			LastModified: time,
			KeyWords:     keyWords,
		}
		practices = append(practices, practice)
	})

	return practices, maxPage, nil
}

func convertToTimestamp(dateStr string) (time.Time, error) {
	layout := "02/01/2006" // Định dạng ngày: dd/MM/yyyy
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
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
