package types

type Config struct {
	Monitor []Monitor `json:"monitor"`
	Daily   []Daily   `json:"daily"`
	Weekly  []Weekly  `json:"weekly"`
	Changed bool      `json:"-"`
}
