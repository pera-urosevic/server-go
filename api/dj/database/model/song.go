package model

import "time"

type Song struct {
	Path     string    `json:"path"`
	Meta     string    `json:"meta"`
	Datetime time.Time `json:"datetime"`
}
