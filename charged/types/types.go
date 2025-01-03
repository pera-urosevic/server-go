package types

type Device struct {
	Name    string `json:"name"`
	Low     int    `json:"low"`
	Value   int    `json:"value"`
	High    int    `json:"high"`
	Command string `json:"command"`
}

type State struct {
	Time    int      `json:"time"`
	Devices []Device `json:"devices"`
}
