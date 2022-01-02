package local

import "github.com/shytikov/advent-of-go/shared"

type Area []*Octopus

func (a *Area) createGrid() (result [][]int) {
	field := *a
	length := len(field)
	x := make([]int, length)
	y := make([]int, length)

	for i, octopus := range field {
		x[i] = octopus.position.X
		y[i] = octopus.position.Y
	}

	lenX := shared.MaxOf(x) + 1
	lenY := shared.MaxOf(y) + 1

	// Create result with correct dimensions
	result = make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		result[i] = make([]int, lenY)
	}

	for _, octopus := range field {
		result[octopus.position.X][octopus.position.Y] = octopus.position.Z
	}

	return
}

func (a *Area) getFlashesCount() (result int) {
	field := *a
	length := len(field)

	for i := 0; i < length; i++ {
		field[i].flashed = false
		if field[i].position.Z == 0 {
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
