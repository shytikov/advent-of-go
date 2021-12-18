package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/05/local"
)

func main() {
	input := local.Read("./input.txt")

	if input.Vents != nil {
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

func solvePuzzleA(input local.Data, result chan int) {
	diagram := input.CreateDiagram()

	for _, vent := range input.Vents {
		if vent.IsOrthogonal() {
			diagram.Draw(vent)
		}
	}

	count := 0

	for i := 0; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] > 1 {
				count++
			}
		}
	}

	result <- count
}

func solvePuzzleB(input local.Data, result chan int) {

	result <- 0
}
