package local

import (
	"github.com/shytikov/advent-of-go/shared"
)

type Node struct {
	root     shared.Point
	children []*Node
}
