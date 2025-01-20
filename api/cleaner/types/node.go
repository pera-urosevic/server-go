package types

type Node struct {
	Entry    Entry  `json:"entry"`
	Children []Node `json:"children"`
}
