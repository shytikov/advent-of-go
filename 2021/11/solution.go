package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/11/local"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntGrid("./input.txt")

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

func solvePuzzleA(input local.Data, result chan int) {
	chain := input.CreateChain()

	fmt.Println(chain)
	fmt.Println(8 >> 1)
	result <- 0
}

func solvePuzzleB(input local.Data, result chan int) {
	result <- 0
}
