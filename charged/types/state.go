package types

type State struct {
	Time    int      `json:"time"`
	Devices []Device `json:"devices"`
}
