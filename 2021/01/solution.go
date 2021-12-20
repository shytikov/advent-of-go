package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"math"
)

func main() {
	input := shared.ReadIntFromLine("./input.txt")

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

func solvePuzzleA(input []int, result chan int) {
	counter := 0
	previous := math.MaxInt16
	for _, current := range input {
		if current-previous > 0 {
			counter++
		}
		previous = current
	}

	result <- counter
}

func solvePuzzleB(input []int, result chan int) {
	counter := 0
	previous := math.MaxInt16
	// It does not matter would we use len(input) or len(input)-2 as missing elements will be replaced with 0
	// and calculation will be still correct, but in this case we will run two unnecessary loop cycles
	for i := 0; i < len(input)-2; i++ {
		current := shared.SumOf(input[i : i+3])
		if current-previous > 0 {
			counter++
		}
		previous = current
	}

	result <- counter
}
