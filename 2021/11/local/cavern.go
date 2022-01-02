package local

import "github.com/shytikov/advent-of-go/shared"

type Cavern []Octopus

func (c Cavern) createGrid() (result [][]int) {
	x := make([]int, len(c))
	y := make([]int, len(c))

	for i, item := range c {
		x[i] = item.position.X
		y[i] = item.position.Y
	}

	lenX := shared.MaxOf(x) + 1
	lenY := shared.MaxOf(y) + 1

	// Create result with correct dimensions
	result = make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		result[i] = make([]int, lenY)
	}

	for _, item := range c {
		result[item.position.X][item.position.Y] = item.position.Z
	}

	return
}

func (c Cavern) getFlashesCount() (result int) {
	for _, link := range c {
		if link.position.Z == 0 {
			link.flashed = false
			result++
		}
	}

	return
}

func (c *Cavern) AccumulateCharge() int {
	// for _, octopus := range *c {
	// }

	return c.getFlashesCount()
}
