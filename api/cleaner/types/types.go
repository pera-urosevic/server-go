package types

type RecordCleaner struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Entry struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Modified string `json:"modified"`
	OK       bool   `json:"ok"`
}
