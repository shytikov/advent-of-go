package local

import "github.com/shytikov/advent-of-go/shared"

type Chain []Node

func (c Chain) createGrid() (result [][]int) {
	x := make([]int, len(c))
	y := make([]int, len(c))

	for i, item := range c {
		x[i] = item.root.X
		y[i] = item.root.Y
	}

	lenX := shared.MaxOf(x) + 1
	lenY := shared.MaxOf(y) + 1

	// Create result with correct dimensions
	result = make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		result[i] = make([]int, lenY)
	}

	for _, item := range c {
		result[item.root.X][item.root.Y] = item.root.Z
	}

	return
}

func (c Chain) SimulateStep() (result int) {
	// for _, link := range c {
	// }

	return
}
