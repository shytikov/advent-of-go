package local

import (
	"sort"

	"github.com/shytikov/advent-of-go/shared"
)

type Data [][]int

func (d Data) CreateArea() (result Area) {
	lenX, lenY := len(d), len(d[0])
	result = make(Area, lenX*lenY)

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			result[d.getOctopusIndex(i, j)] = &shared.Node{
				Value: &Status{position: shared.Point{X: i, Y: j, Z: d[i][j]}, flashed: false},
				Links: []*shared.Node{},
			}
		}
	}

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			indexes := d.getNeighbourIndexes(i, j)
			neighbours := make(Area, len(indexes))

			for k, index := range indexes {
				neighbours[k] = result[index]
			}

			result[d.getOctopusIndex(i, j)].Links = neighbours
		}
	}

	return
}

func (d Data) getOctopusIndex(x, y int) int {
	return x*len(d) + y
}

func (d Data) getNeighbourCount(x, y int) (result int) {
	lenX, lenY := len(d), len(d[0])

	// Maximum possible size of children (neighbors) is 8
	// On an edge it would be 5 (we will use it as 4+1)
	// in the corner it would be 3 (we will use it as 2+1)
	max := 8
	result = max

	// If it's an edge, we exclude 2 neighbors that does not fit
	if x == 0 || x == lenX-1 {
		// First bitwise operation could bring 8 to 4
		result = result >> 1
	}

	// If it's an edge, we exclude 2 neighbors that does not fit
	if y == 0 || y == lenY-1 {
		// Second bitwise opetation could bring it:
		// * either to 4, if it previously was 8
		// * or ot 2, if it previously was 4
		result = result >> 1
	}

	if result != max {
		// If we use bitwise opetations, to get correct result
		// (as we mentioned previously), we need to add 1
		result++
	}

	return
}

func (d Data) getNeighbourIndexes(i, j int) (result []int) {
	result = make([]int, d.getNeighbourCount(i, j))

	lenX, lenY := len(d), len(d[0])

	vector := []shared.Point{
		{X: -1, Y: -1, Z: 0}, // Up-Left
		{X: -1, Y: 0, Z: 0},  // Up-Middle
		{X: -1, Y: 1, Z: 0},  // Up-Right
		{X: 0, Y: -1, Z: 0},  // Center-Left
		{X: 0, Y: 1, Z: 0},   // Center-Right
		{X: 1, Y: -1, Z: 0},  // Down-left
		{X: 1, Y: 0, Z: 0},   // Down-Middle
		{X: 1, Y: 1, Z: 0},   // Down-Right
	}

	count := 0

	for _, direction := range vector {
		x, y := i-direction.X, j-direction.Y
		if x >= 0 && x < lenX && y >= 0 && y < lenY {
			result[count] = d.getOctopusIndex(x, y)
			count++
		}
	}

	sort.Ints(result)

	return
}
