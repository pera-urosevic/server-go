package types

type Weekly struct {
	Task
	LastRun string `json:"lastRun"`
	Time    string `json:"time"`
	Day     string `json:"day"`
}
