package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadIntSlicesFromLines("./input.txt")

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

func solvePuzzleA(input [][]int, result chan int) {
	half := len(input) / 2
	length := len(input[0])

	// Creating a slice to accumulate sums of digits in each position
	accumulator := make([]int, length)
	for _, number := range input {
		for i, digit := range number {
			accumulator[i] = accumulator[i] + digit
		}
	}

	gamma := make([]int, length)
	epsilon := make([]int, length)
	for i, digit := range accumulator {
		if digit > half {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	result <- utils.BinaryIntsToInt(gamma) * utils.BinaryIntsToInt(epsilon)
}

func solvePuzzleB(input [][]int, result chan int) {

	result <- 0
}
