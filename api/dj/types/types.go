package types

import "time"

type QueryResultRaw struct {
	Query    string    `json:"query"`
	Path     string    `json:"path"`
	Meta     string    `json:"meta"`
	Datetime time.Time `json:"datetime"`
}

type QueryResult struct {
	Query    string                 `json:"query"`
	Path     string                 `json:"path"`
	Meta     map[string]interface{} `json:"meta"`
	Datetime time.Time              `json:"datetime"`
}
