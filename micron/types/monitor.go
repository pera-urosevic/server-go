package types

type Monitor struct {
	Task
	Match struct {
		Cmd   string `json:"cmd"`
		Regex string `json:"regex"`
	} `json:"match"`
}
