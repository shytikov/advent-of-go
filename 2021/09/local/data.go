package local

import "github.com/shytikov/advent-of-go/shared"

type Data [][]int

func (d Data) DetectLowPoints() (coordinates []shared.Point2D, heights []int) {
	// Vectors to adjacent locations
	directions := []shared.Point2D{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// Iterating through lines
	for i := 0; i < len(d); i++ {
		// Iterating through columns
		for j := 0; j < len(d[i]); j++ {
			// Flag indicating that we have found the lowest location
			lowest := true
			coordinate, height := shared.Point2D{X: i, Y: j}, d[i][j]

			for _, vector := range directions {
				x, y := i+vector.X, j+vector.Y

				// We're not checking outside the boundaries. Boundaries are walls.
				// We assume our lever is lower than any wall we encounter
				if x >= 0 && y >= 0 && x < len(d) && y < len(d[i]) {
					lowest = lowest && height < d[x][y]
				}
			}

			if lowest {
				coordinates = append(coordinates, coordinate)
				heights = append(heights, height)
			}
		}
	}

	return
}

func (d Data) DetectBasinSizes(coordinates []shared.Point2D) (result []int) {

	return
}
