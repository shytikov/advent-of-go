package shared

type Node struct {
	Value interface{}
	Links []*Node
}