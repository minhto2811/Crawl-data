package crawl

import (
	"fmt"
	"mxgk/crawl/models"
	"mxgk/crawl/service"
	"mxgk/crawl/utils"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var fcmMap map[string]bool = make(map[string]bool)

func CrawlMath() {
	crawlTopics()
	crawlUp()
	crawlHsp()
	crawlGrade()
	crawlSGKTHCS()
	crawlSGKTHPT()
	handleNoti()
}

func handleNoti() {
	//fcmMap["g8"] = true
	for grade, _ := range fcmMap {
		service.SendTopic(grade)
	}
}

func crawlTopics() {
	grade = "g10"
	var source = models.Input{
		Url:   "https://toanmath.com/menh-de-va-tap-hop/page/",
		Type:  "topic",
		Topic: "menhDeTapHop",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/bat-phuong-trinh-bac-nhat-hai-an/page/",
		Type:  "topic",
		Topic: "batPhuongTrinhHaiAn",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/he-phuong-trinh-bac-nhat-ba-an/page/",
		Type:  "topic",
		Topic: "hePhuongTrinhBaAn",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/ham-so-do-thi-va-ung-dung/page/",
		Type:  "topic",
		Topic: "hamSoDoThiUngDung",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/phuong-trinh-he-phuong-trinh-bat-phuong-trinh/page/",
		Type:  "topic",
		Topic: "phuongTrinhHeBatPhuongTrinh",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/bat-dang-thuc-va-cuc-tri/page/",
		Type:  "topic",
		Topic: "batDangThucCucTri",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/he-thuc-luong-trong-tam-giac/page/",
		Type:  "topic",
		Topic: "heThucLuong",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/vecto/page/",
		Type:  "topic",
		Topic: "vecTo",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/thong-ke/page/",
		Type:  "topic",
		Topic: "thongKe",
	}
	autoCrawl(source)

	source = models.Input{
		Url:   "https://toanmath.com/dai-so-to-hop/page/",
		Type:  "topic",
		Topic: "daiSoToHop",
	}
	autoCrawl(source)

	source = models.Input{
		Url:   "https://toanmath.com/xac-suat/page/",
		Type:  "topic",
		Topic: "xacSuat",
	}
	autoCrawl(source)

	source = models.Input{
		Url:   "https://toanmath.com/phuong-phap-toa-do-trong-mat-phang/page/",
		Type:  "topic",
		Topic: "toaDoMatPhang",
	}
	autoCrawl(source)
	grade = "g11"

	source = models.Input{
		Url:   "https://toanmath.com/ham-so-luong-giac-va-phuong-trinh-luong-giac/page/",
		Type:  "topic",
		Topic: "hamSoLuongGiacVaPhuongTrinhLuongGiac",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/day-so-cap-so-cong-va-cap-so-nhan/page/",
		Type:  "topic",
		Topic: "daySoCapSoCongVaCapSoNhan",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/gioi-han-ham-so-lien-tuc/page/",
		Type:  "topic",
		Topic: "gioiHanHamSoLienTuc",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/quan-he-song-song-trong-khong-gian/page/",
		Type:  "topic",
		Topic: "quanHeSongSongTrongKhongGian",
	}
	autoCrawl(source)
	//---------
	source = models.Input{
		Url:   "https://toanmath.com/ham-so-mu-va-ham-so-logarit/page/",
		Type:  "topic",
		Topic: "hamSoMuVaHamSoLogarit",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/dao-ham/page/",
		Type:  "topic",
		Topic: "daoHam",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/quan-he-vuong-goc-trong-khong-gian/page/",
		Type:  "topic",
		Topic: "quanHeVuongGocTrongKhongGian",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/phep-bien-hinh-trong-mat-phang/page/",
		Type:  "topic",
		Topic: "phepBienHinhTrongMatPhang",
	}
	autoCrawl(source)
	grade = "g12"
	source = models.Input{
		Url:   "https://toanmath.com/khao-sat-va-ve-do-thi-ham-so/page/",
		Type:  "topic",
		Topic: "ungDungDaoHamDeKhaoSatVaVeDoThiCuaHamSo",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/nguyen-ham-tich-phan/page/",
		Type:  "topic",
		Topic: "nguyenHamTichPhan",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/phuong-phap-toa-do-trong-khong-gian/page/",
		Type:  "topic",
		Topic: "phuongPhapToaDoTrongKhongGian",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/mat-non-mat-tru-mat-cau/page/",
		Type:  "topic",
		Topic: "matNonMatTruMatCau",
	}
	autoCrawl(source)
	source = models.Input{
		Url:   "https://toanmath.com/so-phuc/page/",
		Type:  "topic",
		Topic: "soPhuc",
	}
	autoCrawl(source)
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
			Url:   "https://toanmath.com/de-thi-thu-thpt-mon-toan/page/",
			Type:  "mock",
			Topic: "",
		},
		{
			Url:   "https://toanmath.com/de-thi-thpt-mon-toan-chinh-thuc/page/",
			Type:  "official",
			Topic: "",
		},
		{
			Url:   "https://toanmath.com/de-danh-gia-nang-luc-mon-toan/page/",
			Type:  "aptitude",
			Topic: "",
		},
		{
			Url:   "https://toanmath.com/tai-lieu-on-thi-thpt-mon-toan/page/",
			Type:  "materials",
			Topic: "",
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
			Url:   "https://thcs.toanmath.com/de-thi-tuyen-sinh-lop-10-mon-toan/page/",
			Type:  "exam",
			Topic: "",
		},
		{
			Url:   "https://thcs.toanmath.com/tai-lieu-toan-on-thi-vao-lop-10/page/",
			Type:  "materials",
			Topic: "",
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
					Url:   fmt.Sprintf("https://thcs.toanmath.com/tai-lieu-toan-%d/page/", g),
					Type:  "materials",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/de-cuong-on-tap-toan-%d/page/", g),
					Type:  "outline",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/de-thi-giua-hk1-toan-%d/page/", g),
					Type:  "midterm1",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/de-thi-hk1-toan-%d/page/", g),
					Type:  "final1",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/de-thi-giua-hk2-toan-%d/page/", g),
					Type:  "midterm2",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/de-thi-hk2-toan-%d/page/", g),
					Type:  "final2",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/khao-sat-chat-luong-toan-%d/page/", g),
					Type:  "assessment",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://thcs.toanmath.com/de-thi-hsg-toan-%d/page/", g),
					Type:  "gifted",
					Topic: "",
				},
			}
			sources = append(sources, s1...)
		} else {
			var s2 = []models.Input{
				{
					Url:   fmt.Sprintf("https://toanmath.com/de-cuong-on-tap-toan-%d/page/", g),
					Type:  "outline",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/de-thi-giua-hk1-toan-%d/page/", g),
					Type:  "midterm1",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/de-thi-hk1-toan-%d/page/", g),
					Type:  "final1",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/de-thi-giua-hk2-toan-%d/page/", g),
					Type:  "midterm2",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/de-thi-hk2-toan-%d/page/", g),
					Type:  "final2",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/khao-sat-chat-luong-toan-%d/page/", g),
					Type:  "assessment",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/de-thi-hsg-toan-%d/page/", g),
					Type:  "gifted",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/giao-an-toan-%d/page/", g),
					Type:  "lessonPlan",
					Topic: "",
				},
				{
					Url:   fmt.Sprintf("https://toanmath.com/tips-giai-toan-%d/page/", g),
					Type:  "tips",
					Topic: "",
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
		p, newPageCount, err, isBreak := getListPractice(url, source.Type, source.Topic)
		if err != nil {
			fmt.Println("err: " + err.Error())
			break
		}
		listPractice = append(listPractice, p...)
		if newPageCount > pageCount {
			pageCount = newPageCount
		}
		if isBreak {
			break
		}
	}

	fmt.Printf("Tài liệu mới: %d\n", len(listPractice))
	if len(listPractice) == 0 {
		return
	}
	fcmMap[listPractice[0].Grade] = true
	for _, practice := range listPractice {
		fmt.Println("Title: " + practice.Title)
		fmt.Println("Grade: " + practice.Grade)
		fmt.Println("Type: " + practice.Type)
		fmt.Println("Topic: " + practice.Topic)
		fmt.Println("Url: " + practice.Url)
		fmt.Println("Last Modified: " + practice.LastModified.Format("02/01/2006"))
		err := updatePracticeToFirestore(practice)
		if err != nil {
			fmt.Println("Lỗi cập nhật lên Firestore:", err)
		}
		fmt.Println("----------------------------")
	}
}

func getListPractice(url string, type1 string, topic string) ([]models.Practice, int, error, bool) {
	fmt.Println("-------------------------")
	doc, err := getDocumentByOption(url)
	if err != nil {
		fmt.Println("Lỗi cào dữ liêu: " + err.Error())
		return nil, 1, err, false
	}
	var practices []models.Practice
	min, _ := convertToTimestamp(minTime)
	max, _ := convertToTimestamp(maxTime)
	maxPage := 1
	doc.Find("a.page-numbers").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}

		re := regexp.MustCompile(`/page/(\d+)`)
		m := re.FindStringSubmatch(href)
		if len(m) == 2 {
			if num, err := strconv.Atoi(m[1]); err == nil && num > maxPage {
				maxPage = num
			}
		}
	})
	isBreak := false
	doc.Find("article").EachWithBreak(func(i int, s *goquery.Selection) bool {
		a := s.Find("h3.entry-title a")
		title := strings.TrimSpace(a.Text())
		link := strings.TrimSpace(a.AttrOr("href", ""))
		b := s.Find("div.mh-meta span.entry-meta-date a")
		dateStr := strings.TrimSpace(b.Text())
		fmt.Printf("> Link tìm thấy ngày %s : %s\n", dateStr, link)
		timeVal, err := convertToTimestamp(dateStr)
		if err != nil {
			fmt.Println("❌ Lỗi convert timestamp:", err)
			return true // dừng luôn nếu lỗi
		}
		if timeVal.Unix() < min.Unix() || timeVal.Unix() >= max.Unix() {
			isBreak = true
			return false // dừng vòng lặp
		}

		doc1, err1 := getDocumentByOption(link)
		if err1 != nil {
			fmt.Println("❌ Error fetching detail page:", err1)
			return true
		}

		// Tìm link PDF đầu tiên
		pdfLink := ""
		doc1.Find("a[href$='.pdf']").EachWithBreak(func(i1 int, s1 *goquery.Selection) bool {
			pdfLink = strings.TrimSpace(s1.AttrOr("href", ""))
			return false // dừng sau khi lấy được link đầu tiên
		})

		if pdfLink == "" {
			fmt.Println("⚠️ Không tìm thấy PDF:", title)
			return true
		}

		withoutDiacritics := utils.RemoveDiacritics(title)
		lowered := strings.ToLower(withoutDiacritics)
		keyWords := strings.Fields(lowered)

		practice := models.Practice{
			Title:        title,
			Grade:        grade,
			Type:         type1,
			Url:          pdfLink,
			LastModified: timeVal,
			KeyWords:     keyWords,
			Topic:        topic,
		}
		practices = append(practices, practice)

		return true // tiếp tục lặp
	})

	return practices, maxPage, nil, isBreak

}

func convertToTimestamp(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	layouts := []string{
		"02/01/2006 15:04:05",
		"02/01/2006 15:04",
		"02/01/2006",
	}
	var t time.Time
	var err error
	for _, l := range layouts {
		t, err = time.Parse(l, dateStr) // time.Parse -> trả UTC
		if err == nil {
			return t.UTC(), nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse time %q", dateStr)
}
