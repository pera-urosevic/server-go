package types

type Task struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Net     bool     `json:"net"`
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
}
