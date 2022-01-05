package local

import "github.com/shytikov/advent-of-go/shared"

type Area []*shared.Node

func (a *Area) createGrid() (result [][]int) {
	area := *a
	length := len(area)
	x := make([]int, length)
	y := make([]int, length)

	for i, current := range area {
		x[i], y[i], _ = (*Octopus)(current).getPosition()
	}

	lenX := shared.MaxOf(x) + 1
	lenY := shared.MaxOf(y) + 1

	// Create result with correct dimensions
	result = make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		result[i] = make([]int, lenY)
	}

	for _, current := range area {
		X, Y, Z := (*Octopus)(current).getPosition()
		result[X][Y] = Z
	}

	return
}

func (a *Area) getFlashesCount() (result int) {
	for _, current := range *a {
		octopus := (*Octopus)(current)
		octopus.setFlashed(false)

		if octopus.getEnergy() == 0 {
			result++
		}
	}

	return
}

func (a *Area) AccumulateCharge() int {
	for _, current := range *a {
		(*Octopus)(current).charge()
	}

	return a.getFlashesCount()
}
