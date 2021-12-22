package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntSlicesFromRuneSlices("./input.txt")

	if input != nil && len(input) > 0 {
		resultA := make(chan int)
		resultB := make(chan int)

		go solvePuzzleA(input, resultA)
		go solvePuzzleB(input, resultB)

		fmt.Println(<-resultA)
		fmt.Println(<-resultB)
	} else {
		panic("failure when reading input data")
	}
}

func solvePuzzleA(input [][]int, result chan int) {
	risk := 0

	// Vectors to adjacent locations
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// Iterating through lines
	for i := 0; i < len(input); i++ {
		// Iterating through columns
		for j := 0; j < len(input[i]); j++ {
			// Flag indicating that we have found the lowest location
			lowest := true
			current := input[i][j]

			for _, vector := range directions {
				x := i + vector[0]
				y := j + vector[1]

				// We're not checking outside the boundaries. Boundaries are walls.
				// We assume our lever is lower than any wall we encounter
				if x >= 0 && y >= 0 && x < len(input) && y < len(input[i]) {
					lowest = lowest && current < input[x][y]
				}
			}

			if lowest {
				risk += current + 1
			}
		}
	}

	result <- risk
}

func solvePuzzleB(input [][]int, result chan int) {
	result <- 0
}
