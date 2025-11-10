package crawl

import (
	"fmt"
	"log"
	"mxgk/crawl/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type InputHocMai struct {
	Url   string
	Grade string
}

const baseURL = "https://hocmai.vn"
const ZCookie = `_fbp=fb.1.1762741483485.114055446582502804; MoodleSession=bm3omum6qc60uss6tcg5el1nk4; MoodleSessionTest=Njkz2RxvgO; _hjSessionUser_239700=eyJpZCI6IjI4ZTYyMzI2LWUwMWEtNTFmNS1hMTM4LTJjM2JjYzJkZGNiYSIsImNyZWF0ZWQiOjE3NjI3NDE1MTkwNTcsImV4aXN0aW5nIjp0cnVlfQ==; G_ENABLED_IDPS=google; g_state={"i_l":0,"i_ll":1762741519572,"i_b":"X5qazY/1DjyAuonk8ElZRSGyj0GUiGPJgPEsFY8lHA0"}; __utmc=267732759; __utmz=267732759.1762741520.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); userchat=%25EE%25D4%2519G%25A7x%25A3%251A%25B0%2516%25C1%2521%25EEU%25D3m%2580%25D0n%252B%25F0%25D0%25C0%2589%2513%2592%25D0E%25C1%25ED%253Fb%2587%2516%25FC%25F9z%25AF%25BB%2595%25C2; MOODLEID_=%25EE%25D4%2519G%25A7x%25A3%251A%25B0%2516%25C1%2521%25EEU%25D3m%2580%25D0n%252B; homeFullname=Giang+Ph%E1%BA%A1m; homeFullname_arr=eyJob21lRnVsbG5hbWUiOiJHaWFuZyBQaFx1MWVhMW0iLCJjaGVja3N1bSI6ImNhMTJiOTQ1YzJiMDdiYmY2ZDdkODhlOWY5NzUzNjMzIn0%3D; __hmrcid=77515715; __utma=267732759.1385023179.1762741519.1762751902.1762754056.3; time_join_home_new=1762754059941; _gid=GA1.2.1775984063.1762754060; _gcl_au=1.1.1799615153.1762754063; _ga_JL7TZSWGMY=GS2.1.s1762754063$o1$g0$t1762754063$j60$l0$h0; _ga=GA1.1.1385023179.1762741519; _tt_enable_cookie=1; _ttp=01K9P58VC6Y4QA19V8JXHKP0GF_.tt.1; ttcsid_CPVP5DJC77UF05LN6OA0=1762754063754::I3Nkxf7KAJEHUpdfuO6g.1.1762754063754.0; ttcsid=1762754063755::S0s-_kMFaKyDiF8qF96H.1.1762754063755.0; __utmt=1; __utmb=267732759.66.10.1762754056; _ga_310661622=GS2.1.s1762754052$o7$g1$t1762756624$j38$l0$h0; _ga_T5W79TC0RG=GS2.1.s1762754052$o3$g1$t1762756625$j37$l0$h0; _ga_2MH9T3CDPY=GS2.1.s1762754052$o12$g1$t1762756625$j37$l0$h0; _ga_659FSX9K28=GS2.1.s1762754052$o12$g1$t1762756625$j37$l0$h0; _ga_YGZBWBRJHV=GS2.1.s1762754052$o12$g1$t1762756625$j37$l0$h0; _ga_2H5R2HXD2Z=GS2.1.s1762754052$o12$g1$t1762756625$j37$l0$h0`
func CrawlHocMai() {
	//crawlLiterature()
	crawlEnglish()
	crawlPhysics()
	crawlBiology()
	crawlHistory()
	crawlChemistry()
	crawlGeography()
	crawlCivics()
}

/*
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "hsp"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=zzz&page=", Grade: "up"},
*/


func crawlCivics() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=278&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=280&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=282&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=192&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=288&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=294&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=300&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=262&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "civics")
	}
}

func crawlGeography() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=264&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=266&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=268&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=190&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=286&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=292&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=298&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=87&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "geography")
	}
}

func crawlChemistry() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=168&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=184&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=270&page=", Grade: "hsp"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=212&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=226&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=240&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=79&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "chemistry")
	}
}

func crawlHistory() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=138&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=156&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=172&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=188&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=272&page=", Grade: "hsp"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=284&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=290&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=296&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=260&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "history")
	}
}


func crawlBiology() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=136&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=154&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=170&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=186&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=214&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=228&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=242&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=80&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "biology")
	}
}

func crawlPhysics() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=134&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=150&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=166&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=182&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=250&page=", Grade: "hsp"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=210&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=224&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=238&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=77&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "physics")
	}
}

func crawlEnglish() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=142&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=148&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=164&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=180&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=200&page=", Grade: "hsp"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=208&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=222&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=236&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=82&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "english")
	}
}

func crawlLiterature() {
	urls := []InputHocMai{
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=130&page=", Grade: "g6"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=144&page=", Grade: "g7"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=160&page=", Grade: "g8"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=176&page=", Grade: "g9"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=198&page=", Grade: "hsp"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=204&page=", Grade: "g10"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=218&page=", Grade: "g11"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=232&page=", Grade: "g12"},
		{Url: "https://hocmai.vn/kho-tai-lieu/list.php?type=category&category=81&page=", Grade: "up"},
	}
	for _, input := range urls {
		grade = input.Grade
		crawl(input.Url, "literatures")
	}
}

func crawl(url string, collection string) {
	var list []models.HocMai
	now := time.Now()
	page := 0
	for i := 0; i <= page; i++ {
		pageURL := url + strconv.Itoa(i)
		fmt.Printf("🚀 Bắt đầu crawl Ngữ văn từ: %s\n", pageURL)
		pageCount, hocmais, err := getList(pageURL, now)
		if err != nil {
			log.Println("Lỗi lấy danh sách tài liệu:", err)
			continue
		}
		list = append(list, hocmais...)
		page = pageCount
		fmt.Printf("Trang %d/%d: Tìm thấy %d tài liệu\n", i+1, page, len(hocmais))
	}
	fmt.Printf("Tổng cộng tìm thấy %d tài liệu Ngữ văn.\n", len(list))

	// Lưu vào Firestore
	for _, hocmai := range list {
		fmt.Println("---------------")
		fmt.Println(hocmai.Title)
		fmt.Println(hocmai.Url)
		fmt.Println(hocmai.LastModified)
		newUrl, err := hmRep.Upload(hocmai.Title, hocmai.Url, ZCookie)
		if err != nil {
			log.Println("Lỗi backup tài liệu:", err)
			continue
		}
		hocmai.Url = newUrl
		err = hmRep.SavePractice(&hocmai, collection)
		if err != nil {
			log.Println("Lỗi lưu tài liệu:", err)
			continue
		}
		fmt.Println("✅ Đã lưu tài liệu:", hocmai.Title)
	}
}

func getList(url string, now time.Time) (int, []models.HocMai, error) {
	// Trang chứa danh sách tài liệu Ngữ văn

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Không thể tải trang:", err)
		return 0, nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Lỗi tải trang: %d %s", res.StatusCode, res.Status)
		return 0, nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Lỗi đọc HTML:", err)
		return 0, nil, err
	}

	var hocmais []models.HocMai

	// Tìm tất cả các thẻ a có class "cap-down"
	doc.Find(".lib-items .item").Each(func(i int, s *goquery.Selection) {
		top := s.Find(".top")
		title := top.Find("a").Text()
		fmt.Println("Tiêu đề:", title)
		download := s.Find("a.cap-down")
		url := download.AttrOr("href", "")
		if strings.HasPrefix(url, "/") {
			url = baseURL + url
		}

		num := -1 * len(hocmais)
		LastModified := now.Add(time.Duration(num) * time.Minute)

		hocmai := models.HocMai{
			Title:        title,
			Url:          url,
			Grade:        grade,
			LastModified: LastModified,
		}
		hocmais = append(hocmais, hocmai)
	})

	if len(hocmais) == 0 {
		fmt.Println("⚠️ Không tìm thấy Tài kiệu nào.")
		return 0, hocmais, nil
	}

	var totalPages int
	doc.Find(".paging a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		// Chỉ lấy các số, bỏ "Tiếp theo"
		if n, err := strconv.Atoi(text); err == nil {
			if n > totalPages {
				totalPages = n
			}
		}
	})
	return totalPages, hocmais, nil
}
