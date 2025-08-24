package crawl

import (
	"fmt"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func CrawlMath() {
	crawlUp()
	crawlHsp()
	crawlGrade()
	crawlSGKTHCS()
	crawlSGKTHPT()
}

func crawlSGKTHCS() {
	grade = "sgk"
	var sources = []models.Input{
		{
			Url:  "https://thcs.toanmath.com/sach-giao-khoa-toan-thcs/page/",
			Type: "thcs",
		},
	}
	for _, source := range sources {
		autoCrawl(source)
	}
}

func crawlSGKTHPT() {
	grade = "sgk"
	var sources = []models.Input{
		{
			Url:  "https://toanmath.com/sach-giao-khoa-toan-thpt/page/",
			Type: "thpt",
		},
	}
	for _, source := range sources {
		autoCrawl(source)
	}
}

func crawlUp() {
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
		autoCrawl(source)
	}
}

func crawlHsp() {
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
		autoCrawl(source)
	}
}

func crawlGrade() {
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
			autoCrawl(source)
		}
	}
}

func autoCrawl(source models.Input) {
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
		updatePracticeToFirestore(practice)
		fmt.Println("----------------------------")
	}
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
		if time.Unix() < min.Unix() || time.Unix() >= max.Unix() {
			fmt.Println("⚠️ Bỏ qua practice:", title)
			return
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

