package local

import "github.com/shytikov/advent-of-go/shared"

type Area []*Octopus

func (a *Area) createGrid() (result [][]int) {
	area := *a
	length := len(area)
	x := make([]int, length)
	y := make([]int, length)

	for i, octopus := range area {
		x[i] = octopus.Value.X
		y[i] = octopus.Value.Y
	}

	lenX := shared.MaxOf(x) + 1
	lenY := shared.MaxOf(y) + 1

	// Create result with correct dimensions
	result = make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		result[i] = make([]int, lenY)
	}

	for _, octopus := range area {
		result[octopus.Value.X][octopus.Value.Y] = octopus.Value.Z
	}

	return
}

func (a *Area) getFlashesCount() (result int) {
	area := *a
	length := len(area)

	for i := 0; i < length; i++ {
		area[i].flashed = false
		if area[i].Value.Z == 0 {
			result++
		}
	}

	return
}

func (a *Area) AccumulateCharge() int {
	for _, octopus := range *a {
		octopus.charge()
	}

	return a.getFlashesCount()
}
