package crawl

import (
	"fmt"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CrawlLiterature() {
	crawl612()
}

func crawl612() {
	for i := 6; i <= 12; i++ {
		grade = fmt.Sprintf("g%d", i)
		g = i
		var sources = []models.Input{
			{
				Url:  fmt.Sprintf("https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-%d", i),
				Type: fmt.Sprintf("g%d", i),
			},
		}
		for _, source := range sources {
			autoCrawlLiterature(source)
		}
	}
}

func autoCrawlLiterature(source models.Input) {
	var listPractice []models.Practice
	pageCount := 1
	for i := 1; i <= pageCount; i++ {
		url := source.Url + fmt.Sprint(i)
		fmt.Println("Đang crawl: " + url)
		p, newPageCount, err := getListLiteratue(url, source.Type)
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
		//UpdatePracticeToFirestore(practice)
		fmt.Println("----------------------------")
	}
}

func getListLiteratue(url string, type1 string) ([]models.Practice, int, error) {
	doc, err := getDocument(url)
	if err != nil {
		return nil, 1, err
	}

	var practices []models.Practice
	min, _ := convertToTimestamp(minTime)
	max, _ := convertToTimestamp(maxTime)
	maxPage := 1
	doc.Find(".page-nav").Each(func(i int, s *goquery.Selection) {
		currentNode := s.Find(".current").First()
		currentStr  := strings.TrimSpace(currentNode.Text())
		current, err := strconv.Atoi(currentStr)
		if err == nil && current > maxPage {
			maxPage = current
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
