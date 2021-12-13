package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/05/local"
)

func main() {
	input := local.Read("./input.txt")

	if input.Vents != nil {

		fmt.Println(input)
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

	result <- 0
}

func solvePuzzleB(input local.Data, result chan int) {

	result <- 0
}
