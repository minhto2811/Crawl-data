package crawl

import (
	"context"
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
	"github.com/go-rod/rod/lib/proto"
)

const (
	minTime = "16/12/2025"
	maxTime = "20/12/2025"
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
    // 1. Khởi tạo browser với các tùy chọn stealth
    launcherArgs := []string{
        "--no-sandbox",
        "--disable-setuid-sandbox",
        "--disable-blink-features=AutomationControlled",
        "--disable-dev-shm-usage",
        "--disable-web-security",
        "--disable-features=IsolateOrigins,site-per-process",
        "--disable-site-isolation-trials",
        "--disable-accelerated-2d-canvas",
        "--disable-gpu",
        "--window-size=1920,1080",
    }
    
    // Tạo launcher với các arguments
    l := launcher.New().
        Headless(true). // Đặt true để chạy ẩn
        Devtools(false).
        Set("disable-blink-features", "AutomationControlled").
        UserDataDir("rod_browser_data")
    
    // Thêm các arguments
    for _, arg := range launcherArgs {
    l = l.Append("arg", arg)
}
    
    // Khởi động browser
    controlURL, err := l.Launch()
    if err != nil {
        return nil, fmt.Errorf("không thể khởi động browser: %v", err)
    }
    
    // Kết nối tới browser
    browser := rod.New().ControlURL(controlURL)
    
    // Kết nối với timeout
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    browser = browser.Context(ctx)
    err = browser.Connect()
    if err != nil {
        return nil, fmt.Errorf("không thể kết nối browser: %v", err)
    }
    defer func() {
        // Đảm bảo browser đóng ngay cả khi có panic
        if r := recover(); r != nil {
            _ = browser.Close()
            panic(r)
        }
    }()
    
    // 2. Tạo page với stealth mode
    page, err := browser.Page(proto.TargetCreateTarget{})
    if err != nil {
        return nil, fmt.Errorf("không thể tạo page: %v", err)
    }
    
    // Thiết lập User-Agent giống trình duyệt thật
    err = page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
        UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
        AcceptLanguage: "vi-VN,vi;q=0.9,en-US;q=0.8,en;q=0.7",
        Platform:       "Windows",
    })
    if err != nil {
        return nil, fmt.Errorf("không thể thiết lập user agent: %v", err)
    }
    
    // Thêm các headers khác
    extraHeaders := map[string]string{
        "Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
        "Accept-Language": "vi-VN,vi;q=0.9,en-US;q=0.8,en;q=0.7",
        "Accept-Encoding": "gzip, deflate, br",
        "Connection":      "keep-alive",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest": "document",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-Site": "none",
        "Cache-Control":   "max-age=0",
    }
    
    for key, value := range extraHeaders {
        page.MustSetExtraHeaders(key, value)
    }
    
    // 3. Thực hiện JavaScript stealth để tránh bị phát hiện
    stealthJS := `
    () => {
        // Xóa webdriver flag
        delete navigator.__proto__.webdriver;
        
        // Thay đổi plugins
        Object.defineProperty(navigator, 'plugins', {
            get: () => [1, 2, 3, 4, 5],
        });
        
        // Thay đổi languages
        Object.defineProperty(navigator, 'languages', {
            get: () => ['vi-VN', 'vi', 'en-US', 'en'],
        });
        
        // Thay đổi platform
        Object.defineProperty(navigator, 'platform', {
            get: () => 'Win32',
        });
        
        // Mock Chrome runtime
        window.chrome = {
            runtime: {},
        };
        
        // Ẩn permission query
        const originalQuery = window.navigator.permissions.query;
        window.navigator.permissions.query = (parameters) => (
            parameters.name === 'notifications' ?
                Promise.resolve({ state: Notification.permission }) :
                originalQuery(parameters)
        );
    }
    `
    
    page.MustEval(stealthJS)
    
    // 4. Điều hướng đến URL
    err = page.Navigate(url)
    if err != nil {
        return nil, fmt.Errorf("không thể điều hướng đến URL: %v", err)
    }
    
    // 5. Chờ page load và xử lý Cloudflare
    page.MustWaitLoad()
    
    // Chờ cho các request hoàn tất
    page.Timeout(30 * time.Second).MustWaitRequestIdle()
    
    // Kiểm tra Cloudflare challenge
    var cloudflareDetected bool
    maxRetries := 3
    retryDelay := 5 * time.Second
    
    for i := 0; i < maxRetries; i++ {
        html, err := page.HTML()
        if err != nil {
            return nil, fmt.Errorf("không thể lấy HTML: %v", err)
        }
        
        // Kiểm tra các dấu hiệu Cloudflare
        cloudflareDetected = strings.Contains(html, "cf-browser-verification") ||
            strings.Contains(html, "challenge-form") ||
            strings.Contains(html, "cf_chl_opt") ||
            strings.Contains(html, "cf_captcha") ||
            strings.Contains(html, "cloudflare")
        
        if !cloudflareDetected {
            break
        }
        
        fmt.Printf("Phát hiện Cloudflare (lần thử %d/%d)...\n", i+1, maxRetries)
        
        if i < maxRetries-1 {
            // Thử xử lý challenge
            handleCloudflareChallenge(page)
            time.Sleep(retryDelay)
        }
    }
    
    if cloudflareDetected {
        return nil, fmt.Errorf("không thể vượt qua Cloudflare sau %d lần thử", maxRetries)
    }
    
    // 6. Lấy HTML cuối cùng
    html, err := page.HTML()
    if err != nil {
        return nil, fmt.Errorf("không thể lấy HTML cuối cùng: %v", err)
    }
    
    // 7. Chuyển đổi sang goquery Document
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        return nil, fmt.Errorf("không thể parse HTML: %v", err)
    }
    
    return doc, nil
}

func handleCloudflareChallenge(page *rod.Page) {
    // Thử các cách khác nhau để xử lý challenge
    
    // Cách 1: Thử click nút submit nếu có
    if el, err := page.Element("input[type='submit'], button[type='submit']"); err == nil {
        el.MustClick()
        page.MustWaitRequestIdle()
        time.Sleep(2 * time.Second)
        return
    }
    
    // Cách 2: Thử click div có class chứa "cf"
    if el, err := page.Element("div[class*='cf'], div[class*='cloudflare']"); err == nil {
        el.MustClick()
        page.MustWaitRequestIdle()
        time.Sleep(2 * time.Second)
        return
    }
    
    // Cách 3: Refresh page
    page.MustReload()
    page.MustWaitLoad()
    time.Sleep(3 * time.Second)
}

// Hàm helper để đóng browser an toàn
func cleanupBrowser(browser *rod.Browser) {
    if browser != nil {
        _ = browser.Close()
    }
}