package model

func (Query) TableName() string {
	return "queries"
}

type Query struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}
