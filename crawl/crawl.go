package crawl

import (
	"encoding/json"
	"fmt"
	"log"
	"mxgk/crawl/models"
	"mxgk/crawl/repo"
	"net/http"
	"strings"
	"time"

	"os/exec"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	minTime = "20/10/2025"
	maxTime = "21/10/2025" //Nếu bằng ngày hôm nay tức là đã crawl
)

var grade = "g8"
var g = 8
var collection = "practices"

var rep repo.PracticeRepo

func SetRepo(practiceRepo repo.PracticeRepo) {
	rep = practiceRepo
}

type VideoItem struct {
	ID string `json:"id"`
}

func bytesToLines(b []byte) []string {
	lines := []string{}
	line := []byte{}
	for _, c := range b {
		if c == '\n' {
			lines = append(lines, string(line))
			line = []byte{}
		} else {
			line = append(line, c)
		}
	}
	if len(line) > 0 {
		lines = append(lines, string(line))
	}
	return lines
}

func CrawlVideo(playlist string,gr string) {
cmd := exec.Command("yt-dlp", "-j", "--flat-playlist", playlist)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	// Mỗi dòng là một JSON object -> parse từng dòng
	lines := bytesToLines(output)
	var urls []string 
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var item VideoItem
		if err := json.Unmarshal([]byte(line), &item); err == nil && item.ID != "" {
			url := "https://www.youtube.com/watch?v=" + item.ID
			urls = append(urls, url)
			fmt.Println(url)
		}
	}

	var videos []models.Video
    for _, url := range urls {
        title, err := getYouTubeTitle(url)
        if err != nil {
            fmt.Println("Lỗi lấy title:", err)
            continue
        }
		if title == "- YouTube" {
            fmt.Println("Xóa video pivate")
            continue
        }
        videos = append(videos, models.Video{
            Title:        title,
            URL:          url,
            Grade:        gr,
            LastModified: time.Now(),
			Playlist: playlist,
        })
    }
    rep.Tutorial(videos)
}

func getYouTubeTitle(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    tokenizer := html.NewTokenizer(resp.Body)
    for {
        tt := tokenizer.Next()
        switch tt {
        case html.ErrorToken:
            return "", fmt.Errorf("title not found")
        case html.StartTagToken:
            t := tokenizer.Token()
            if t.Data == "title" {
                tokenizer.Next()
                title := strings.TrimSpace(tokenizer.Token().Data)
                // YouTube title thường có “ - YouTube”, ta bỏ đi
                return strings.TrimSuffix(title, " - YouTube"), nil
            }
        }
    }
}



func BackUp() error {
	return rep.Backup()
}

func updatePracticeToFirestore(practice models.Practice) error {
	err := rep.SavePractice(&practice, collection)
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
