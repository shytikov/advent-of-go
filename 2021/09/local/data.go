package local

import (
	"fmt"
	"sync"

	"github.com/shytikov/advent-of-go/shared"
)

type Data [][]int

// Vectors to adjacent locations
var directions = []shared.Point2D{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func (d Data) DetectLowestPoints() (coordinates []shared.Point2D, heights []int) {
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
	result = make([]int, len(coordinates))

	var wg sync.WaitGroup

	for i, coordinate := range coordinates {
		wg.Add(1)

		go func(index int, lowest shared.Point2D) {
			defer wg.Done()

			queue := []shared.Point2D{lowest}
			count := len(queue)

			visited := map[string]bool{
				fmt.Sprintf("%v,%v", lowest.X, lowest.Y): true,
			}

			for j := 0; j < len(queue); j++ {
				point := queue[j]

				for _, vector := range directions {
					x, y := point.X+vector.X, point.Y+vector.Y

					// We're not checking outside the boundaries. Boundaries are walls.
					if x >= 0 && y >= 0 && x < len(d) && y < len(d[0]) {
						current := fmt.Sprintf("%v,%v", x, y)
						_, found := visited[current]

						if d[x][y] < 9 && !found {
							queue = append(queue, shared.Point2D{x, y})
							visited[current] = true
							count++
						}
					}
				}
			}

			result[index] = count
		}(i, coordinate)
	}

	wg.Wait()

	return
}
