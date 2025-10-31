package models

type Input struct {
	Url  string `json:"url"`  // URL của bài tập
	Type string `json:"type"` // Loại bài tập (ví dụ: "Toán", "Văn", v.v.)
	Topic string `json:"topic,omitempty"` // Chủ đề cụ thể (nếu có)
}
