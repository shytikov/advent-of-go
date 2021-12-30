package local

import (
	"sort"

	"github.com/shytikov/advent-of-go/shared"
)

type Data [][]int

func (d Data) CreateChain() (result Chain) {
	lenX, lenY := len(d), len(d[0])
	result = make(Chain, lenX*lenY)

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			root := shared.Point{i, j, d[i][j]}
			children := make([]*Node, d.getChildrenCount(i, j))

			result[d.getLinkIndex(i, j)] = Node{root, children}
		}
	}

	result.assembleChildren(lenX, lenY)

	return
}

func (d Data) getLinkIndex(x, y int) int {
	return x*len(d) + y
}

func (d Data) getChildrenCount(x, y int) (result int) {
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

func (d Data) getChildrenIndexes(i, j int) (result []int) {
	result = make([]int, d.getChildrenCount(i, j))

	lenX, lenY := len(d), len(d[0])

	vector := []shared.Point{
		{1, 0, 0},
		{1, 1, 0},
		{0, 1, 0},
		{0, -1, 0},
		{-1, 1, 0},
		{-1, 0, 0},
		{-1, -1, 0},
	}

	count := 0

	for _, direction := range vector {
		x, y := i-direction.X, j-direction.Y
		if x >= 0 && x < lenX && y >= 0 && y < lenY {
			result[count] = d.getLinkIndex(x, y)
			count++
		}
	}

	sort.Ints(result)

	return
}
