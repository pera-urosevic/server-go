package model

func (Bookmark) TableName() string {
	return "bookmarks"
}

type Bookmark struct {
	Path string `json:"path"`
}
