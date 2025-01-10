package types

type Photo struct {
	ID          int64  `json:"id"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	Modified    int64  `json:"modified"`
	Online      bool   `json:"online"`
	Album       string `json:"album"`
	Datetime    string `json:"datetime"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Copyright   string `json:"copyright"`
}
