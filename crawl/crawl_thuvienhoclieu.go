package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"mxgk/crawl/models"
	"mxgk/crawl/utils"
	"strconv"
	"strings"
	"time"
)

func ClearTVHL(cutoff time.Time) error {
	return tvhlRep.Clear(cutoff)
}

func CrawlTVHL() {
	utils.CreateClient()
	crawlLiteratureTVHL()
	crawlNaturalScienceTVHL()
	crawlEnglishTVHL()
	crawlPhysicsTVHL()
	crawlBiologyTVHL()
	crawlHistoryTVHL()
	crawlChemistryTVHL()
	crawlGeographyTVHL()
	 crawlCivicsTVHL()
}

func crawlNaturalScienceTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-6/khtn-lop-6-canh-dieu/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-6/khtn-lop-6-chan-troi-sang-tao/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-6/khtn-lop-6-ket-noi-tri-thuc/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-7/khtn-7-ctst/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-7/khtn-7-canh-dieu/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-7/khtn-7-kntt/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-7/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-khoa-hoc-tu-nhien/khtn-lop-9/", Grade: "g9"},
	}
	for _, input := range urls {
		crawlTVHL(input.Url, input.Grade, "naturalScience")
	}
}

func crawlCivicsTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-6/gdcd-6-sach-canh-dieu/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-6/gdcd-6-sach-ket-noi-tri-thuc/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-6/gdcd-6-sach-chan-troi-sang-tao/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-7/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-gdcd-lop-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-cong-dan-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-cong-dan-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-cong-dan-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-cong-dan/tai-lieu-cong-dan-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "civics")
	}
}

func crawlGeographyTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-9-tai-lieu-dia-li/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-dia-li/tai-lieu-dia-li-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "geography")
	}
}

func crawlChemistryTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-hoa-hoc/tai-lieu-hoa-hoc-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-hoa-hoc/tai-lieu-hoa-hoc-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-hoa-hoc/tai-lieu-hoa-hoc-lop-12/de-kiem-tra-giua-hoc-ky-1-hoa-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-hoa-hoc/tai-lieu-hoa-hoc-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-hoa-hoc/tai-lieu-hoa-hoc-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "chemistry")
	}
}

func crawlHistoryTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-7/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-lich-su/tai-lieu-lich-su-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "history")
	}
}

func crawlBiologyTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-sinh-hoc/tai-lieu-sinh-hoc-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-sinh-hoc/tai-lieu-sinh-hoc-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-sinh-hoc/tai-lieu-sinh-hoc-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-sinh-hoc/tai-lieu-sinh-hoc-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "biology")
	}
}

func crawlPhysicsTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-vat-li/tai-lieu-vat-li-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-vat-li/tai-lieu-vat-li-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-vat-li/tai-lieu-vat-li-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-vat-li/tai-lieu-vat-li-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "physics")
	}
}

func crawlEnglishTVHL() {
	urls := []models.InputTVHL{
		//{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-6/de-kiem-tra-giua-hoc-ky-2-tieng-anh-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-6/tieng-anh-6-kntt/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-7/de-kiem-tra-giua-hoc-ky-2-tieng-anh-7/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-7/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-8/de-kiem-tra-giua-hoc-ky-2-tieng-anh-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-9/de-kiem-tra-giua-hoc-ky-1-tieng-anh-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-9/de-kiem-tra-giua-hoc-ky-2-tieng-anh-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-tieng-anh/tai-lieu-tieng-anh-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {

		crawlTVHL(input.Url, input.Grade, "english")
	}
}

func crawlLiteratureTVHL() {
	urls := []models.InputTVHL{
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-6/ngu-van-6-sach-canh-dieu/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-6/ngu-van-6-sach-ket-noi-tri-thuc/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-6/ngu-van-6-sach-chan-troi-sang-tao/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-6/", Grade: "g6"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-7/", Grade: "g7"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-8/", Grade: "g8"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-9/", Grade: "g9"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-10/", Grade: "g10"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-11/", Grade: "g11"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/tai-lieu-ngu-van-lop-12/", Grade: "g12"},
		{Url: "https://thuvienhoclieu.com/tai-lieu-ngu-van/ngu-van-luyen-thi/", Grade: "up"},
	}
	for _, input := range urls {
		crawlTVHL(input.Url, input.Grade, "literatures")
	}
}

func crawlTVHL(url string, grade1 string, collection string) {
	fmt.Printf("🚀 Bắt đầu crawlTVHL từ: %s\n", url)
	var list []models.TVHL
	page := 1
	for i := 1; i <= page; i++ {
		fmt.Printf("Trang số: %d\n", i)
		var pageURL string
		if i == 1 {
			pageURL = url
		} else {
			pageURL = url + "page/" + strconv.Itoa(i)
		}

		pageCount, hocmais, err := getListTVHL(pageURL, grade1)
		if err != nil {
			fmt.Println("Lỗi lấy danh sách tài liệu:", err)
			continue
		}
		list = append(list, hocmais...)
		page = pageCount
		if(page > 10) {
			page = 10
		}
		fmt.Printf("Thêm %d tài liệu\n", len(hocmais))
	}

	total := len(list)
	fmt.Printf("Tổng cộng tìm thấy %d tài liệu.\n", total)

	//Lưu vào Firestore
	for i, hocmai := range list {
		fmt.Printf("------------------ %d/%d ------------------\n", i+1, total)
		newUrl, err := uploadAndConvert(hocmai.Title, hocmai.Url, i%2)
		if err != nil {
			log.Println("Lỗi backup tài liệu:", err)
			continue
		}

		hocmai.Url = newUrl
		if err := tvhlRep.SavePractice(&hocmai, collection); err != nil {
			log.Println("Lỗi lưu tài liệu:", err)
			continue
		}

		fmt.Println("✅ Đã lưu tài liệu:", hocmai.Title)
	}
}

var loSem = make(chan int, 2)

func uploadAndConvert(title, url string, workerID int) (string, error) {
	loSem <- workerID
	defer func() { <-loSem }()

	// Upload sẽ tải file + gọi ConvertDocxToPDF bên trong
	return tvhlRep.Upload(title, url)
}

func getListTVHL(url string, grade1 string) (int, []models.TVHL, error) {
	res, err := utils.CreateRequest(url, "GET")
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return 0, nil, fmt.Errorf("HTTP %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Lỗi đọc HTML:", err)
		return 0, nil, err
	}

	var hocmais []models.TVHL

	sel := doc.Find(".td-block-span6 .td_module_1")
	fmt.Println("Count:", sel.Length())
	finish := false
	sel.Each(func(i int, s *goquery.Selection) {
		if finish {
			return
		}
		a := s.Find(".entry-title a")
		link, exists := a.Attr("href")
		if !exists {
			log.Println("Không tìm thấy link tài liệu")
			return
		}

		timeNode := s.Find(".td-post-date time")
		dateStr, ok := timeNode.Attr("datetime")
		if !ok {
			log.Println("Không tìm thấy ngày tháng")
			return
		}

		date, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			log.Println("Lỗi phân tích ngày tháng:", err)
			return
		}

		min, _ := utils.ConvertToTimestamp(minTime)
		max, _ := utils.ConvertToTimestamp(maxTime)

		if date.Unix() < min.Unix() || date.Unix() >= max.Unix() {
			log.Printf("Bỏ qua tài liệu '%s' vì không trong khoảng thời gian.\n", a.Text())
			finish = true
			return
		}

		linkDownload, title, err := getLinkDowloadTVHL(link)
		if err != nil {
			log.Println("Lỗi lấy link download:", err)
			return
		}

		fmt.Printf(">>> %s\n", title)

		hocmai := models.TVHL{
			Title:        title,
			Grade:        grade1,
			Url:          linkDownload,
			LastModified: date,
		}
		hocmais = append(hocmais, hocmai)
	})

	// Tìm tất cả các thẻ a có class "cap-down"

	if len(hocmais) == 0 {
		fmt.Println("⚠️ Không tìm thấy Tài kiệu nào.")
		return 0, hocmais, nil
	}

	lastText := strings.TrimSpace(
		doc.Find(".page-nav a.last").Text(),
	)

	totalPages, _ := strconv.Atoi(lastText)

	return totalPages, hocmais, nil
}

func getLinkDowloadTVHL(url string) (string, string, error) {
	res, err := utils.CreateRequest(url, "GET")
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", fmt.Errorf("HTTP %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Lỗi đọc HTML:", err)
		return "", "", err
	}

	a := doc.Find("p.embed_download a")

	downloadLink, exists := a.Attr("href")
	if !exists {
		log.Println("❌ Không tìm thấy link download")
		return "", "", fmt.Errorf("Không tìm thấy link download")
	}
	title := doc.Find("header.td-post-title h1.entry-title").Text()
	return downloadLink, title, nil

}
