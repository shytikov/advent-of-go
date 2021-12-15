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
		if vent.IsHorizontal() {
			var begin int
			var end int

			if vent.From.X > vent.To.X {
				begin = vent.To.X
				end = vent.From.X
			} else {
				begin = vent.From.X
				end = vent.To.X
			}

			y := vent.From.Y

			for i := begin; i <= end; i++ {
				diagram[i][y] = diagram[i][y] + 1
			}
		}

		if vent.IsVertical() {
			var begin int
			var end int

			if vent.From.Y > vent.To.Y {
				begin = vent.To.Y
				end = vent.From.Y
			} else {
				begin = vent.From.Y
				end = vent.To.Y
			}

			x := vent.From.X

			for i := begin; i <= end; i++ {
				diagram[x][i] = diagram[x][i] + 1
			}
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
