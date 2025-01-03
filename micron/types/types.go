package types

type Task struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Net     bool     `json:"net"`
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
}

type Monitor struct {
	Task
	Match struct {
		Cmd   string `json:"cmd"`
		Regex string `json:"regex"`
	} `json:"match"`
}

type Daily struct {
	Task
	LastRun string `json:"lastRun"`
	Time    string `json:"time"`
}

type Weekly struct {
	Task
	LastRun string `json:"lastRun"`
	Time    string `json:"time"`
	Day     string `json:"day"`
}

type Config struct {
	Monitor []Monitor `json:"monitor"`
	Daily   []Daily   `json:"daily"`
	Weekly  []Weekly  `json:"weekly"`
	Changed bool      `json:"-"`
}

type Status struct {
	Enabled bool `json:"enabled"`
}
