package types

type Daily struct {
	Task
	LastRun string `json:"lastRun"`
	Time    string `json:"time"`
}
