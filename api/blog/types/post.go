package types

type Post struct {
	ID          int64  `json:"id"`
	Timestamp   string `json:"timestamp"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Template    string `json:"template"`
	Description string `json:"description"`
	Image       string `json:"image"`
	URL         string `json:"url"`
}
