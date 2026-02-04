package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var client *http.Client
var targetURL = "https://thuvienhoclieu.com/"

func CreateClient() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   15 * time.Second,
			KeepAlive: 15 * time.Second,
			DualStack: false,
		}).DialContext,
		ForceAttemptHTTP2: false,
	}

	client = &http.Client{
		Transport: transport,
		Timeout:   20 * time.Second,
	}
	jar, _ := cookiejar.New(nil)
	client.Jar = jar

	// request trang chủ trước
	res, err := CreateRequest(targetURL, "GET")
	if err != nil {
		println("Lấy cookie thất bại :)")
		return
	}
	defer res.Body.Close()
	io.Copy(io.Discard, res.Body)
}

func CreateRequest(url string, method string) (*http.Response, error) {
	sleep := time.Duration(2) * time.Second
	fmt.Println("Bắt đầu sau 2s")
	time.Sleep(sleep)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/120.0.0.0 Safari/537.36")

	req.Header.Set("Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")

	req.Header.Set("Accept-Language", "vi-VN,vi;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Referer", "https://thuvienhoclieu.com/")
	res, err := client.Do(req)
	return res, err
}
