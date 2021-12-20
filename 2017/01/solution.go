package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadIntSlicesFromRuneSlices("./input.txt")

	if input != nil {
		resultA := make(chan int)
		resultB := make(chan int)

		go solvePuzzleA(input[0], resultA)
		go solvePuzzleB(input[0], resultB)

		fmt.Println(<-resultA)
		fmt.Println(<-resultB)
	} else {
		panic("failure when reading input data")
	}
}

func solvePuzzleA(input []int, result chan int) {
	input = append(input, input[0])
	accumulator := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			accumulator += input[i]
		}
	}

	result <- accumulator
}

func solvePuzzleB(input []int, result chan int) {
	input = append(input, input[0:len(input)/2]...)

	result <- 0
}
