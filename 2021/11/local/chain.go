package local

import (
	"github.com/shytikov/advent-of-go/shared"
)

type Chain []Node

func (c Chain) assembleChildren(lenX, lenY int) {
	directions := []shared.Point{
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 1},
		{X: -1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: -1},
	}

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			//root := c[i*(lenX-1)+j]
			//size := root.getChildrenCount(lenX, lenY)
			//root.children = make([]*Node, size)

			count := 0
			for _, coordinate := range directions {
				x, y := i+coordinate.X, j+coordinate.Y
				if x != -1 && x != lenX && y != -1 && y != lenY {
					//fmt.Println(i, j, x, y, size, count, coordinate, x*(lenX-1)+y)
					//root.children[count] = &c[x*(lenX-1)+y]
					count++
				}
			}
		}
	}
}
