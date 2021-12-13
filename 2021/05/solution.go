package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/05/local"
)

const dimension = 5

func main() {
	input := local.Read("./input.txt")

	if input != nil {
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

func solvePuzzleA(input []local.Vent, result chan int) {

	result <- 0
}

func solvePuzzleB(input []local.Vent, result chan int) {

	result <- 0
}
