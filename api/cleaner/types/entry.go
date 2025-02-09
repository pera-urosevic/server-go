package types

type Entry struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Modified string `json:"modified"`
	Status   int64  `json:"status"`
}
