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
	result <- 0
}

func solvePuzzleB(input [][]int, result chan int) {
	result <- 0
}
