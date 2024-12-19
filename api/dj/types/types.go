package types

import "time"

type RecordSong struct {
	Path     string                 `json:"path"`
	Meta     map[string]interface{} `json:"meta"`
	Datetime time.Time              `json:"datetime"`
}

type RecordQuery struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type RecordQueryResult struct {
	Query    string                 `json:"query"`
	Path     string                 `json:"path"`
	Meta     map[string]interface{} `json:"meta"`
	Datetime time.Time              `json:"datetime"`
}
