package local

import (
	"fmt"
	"sync"

	"github.com/shytikov/advent-of-go/shared"
)

type Data [][]int

// Vectors to adjacent locations
var directions = []shared.Point{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

func (d Data) DetectLowestPoints() (result []shared.Point) {
	// Iterating through lines
	for i := 0; i < len(d); i++ {
		// Iterating through columns
		for j := 0; j < len(d[i]); j++ {
			// Flag indicating that we have found the lowest location
			lowest := true
			coordinate := shared.Point{
				X: i,
				Y: j,
				Z: d[i][j],
			}

			for _, vector := range directions {
				x, y := i+vector.X, j+vector.Y

				// We're not checking outside the boundaries. Boundaries are walls.
				// We assume our lever is lower than any wall we encounter
				if x >= 0 && y >= 0 && x < len(d) && y < len(d[i]) {
					lowest = lowest && coordinate.Z < d[x][y]
				}
			}

			if lowest {
				result = append(result, coordinate)
			}
		}
	}

	return
}

func (d Data) DetectBasinSizes(coordinates []shared.Point) (result []int) {
	result = make([]int, len(coordinates))

	var wg sync.WaitGroup

	for i, coordinate := range coordinates {
		wg.Add(1)

		go func(index int, lowest shared.Point) {
			defer wg.Done()

			queue := []shared.Point{lowest}
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
							queue = append(queue, shared.Point{x, y, d[x][y]})
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
