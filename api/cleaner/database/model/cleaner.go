package model

func (Cleaner) TableName() string {
	return "cleaner"
}

type Cleaner struct {
	Path string `json:"path"`
	Name string `json:"name"`
}
