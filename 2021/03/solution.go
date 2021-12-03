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

	// Gamma – most common digit
	γ := make([]int, length)
	// Epsilon – least common digit
	ε := make([]int, length)
	for i, digit := range accumulator {
		if digit > half {
			γ[i] = 1
			ε[i] = 0
		} else {
			γ[i] = 0
			ε[i] = 1
		}
	}

	result <- utils.BinaryIntsToInt(γ) * utils.BinaryIntsToInt(ε)
}

func solvePuzzleB(input [][]int, result chan int) {

	result <- 0
}
