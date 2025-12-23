package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2/google"
)

type FCMNotification struct {
	Title string            `json:"title"`
	Body  string            `json:"body"`
	Data  map[string]string `json:"data,omitempty"`
}

type FCMService struct {
	ProjectID  string
	SAFilePath string
}

// Lấy access token từ service account JSON
func (f *FCMService) getAccessToken() (string, error) {
	data, err := os.ReadFile(f.SAFilePath)
	if err != nil {
		return "", err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		return "", err
	}
	token, err := conf.TokenSource(context.Background()).Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

// Gửi notification đến topic
func (f *FCMService) SendToTopic(topic string, notification FCMNotification) error {
	token, err := f.getAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://fcm.googleapis.com/v1/projects/%s/messages:send", f.ProjectID)

	message := map[string]interface{}{
		"message": map[string]interface{}{
			"topic": topic,
			"notification": map[string]string{
				"title": notification.Title,
				"body":  notification.Body,
			},
		},
	}

	if len(notification.Data) > 0 {
		message["message"].(map[string]interface{})["data"] = notification.Data
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("FCM send error: %s, response: %s", resp.Status, string(respBody))
	}

	fmt.Printf("Notification sent to topic '%s' successfully!\n", topic)
	return nil
}

func SendTopic(grade string)  {
	projectID := "dehay-73822"
	rootPath, _ := os.Getwd()
	serviceAccount := filepath.Join(rootPath, "serviceAccountKey.json")
	fcm := FCMService{
		ProjectID:  projectID,
		SAFilePath: serviceAccount,
	}
	notification := FCMNotification{
		Title: "Có tài liệu mới được cập nhật",
		Body:  "Chúng mình vừa cập nhật tài liệu, vào ứng dụng để xem nhé. Chúc bạn học tập vui vẻ!",
		Data:  map[string]string{},
	}
	fcm.SendToTopic(grade, notification)
	
}
