package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadNumbersFromString("./input.txt")

	if input != nil {
		resultA := make(chan int)
		resultB := make(chan int)

		go solvePuzzleA(input, resultA)
		go solvePuzzleB(input, resultB)

		fmt.Println(<-resultA)
		fmt.Println(<-resultB)
	} else {
		fmt.Errorf("failure when reading input data")
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
