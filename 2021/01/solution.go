package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
	"math"
)

func main() {
	input := utils.ReadNumbers("./input.txt")

	if input != nil {
		// In this approach we will be reading input slice from two concurrent goroutines which is not ideal,
		// but as soon as we are not modifying it, no issues should occur 🤞

		// Channels for respective puzzles' solution
		firstResult := make(chan int)
		secondResult := make(chan int)

		go solveFirstPuzzle(input, firstResult)
		go solveSecondPuzzle(input, secondResult)

		fmt.Println(<-firstResult)
		fmt.Println(<-secondResult)
	} else {
		fmt.Errorf("failure to read input data")
	}
}

func solveFirstPuzzle(input []int, result chan int) {
	var counter int

	previous := math.MaxInt16
	for _, current := range input {
		previous, counter = decide(current, previous, counter)
	}

	result <- counter
}

func solveSecondPuzzle(input []int, result chan int) {
	var counter int

	previous := math.MaxInt16
	// It does not matter would we use len(input) or len(input)-2 as missing elements will be replaced with 0
	// and calculation will be still correct, but in this case we will run two unnecessary loop cycles
	for i := 0; i < len(input)-2; i++ {
		previous, counter = decide(utils.SumNumbers(input[i:i+3]), previous, counter)
	}

	result <- counter
}

// Function will decide which value should be used for counter
// In addition it will help to swap current and previous values in one statement
func decide(current, previous, counter int) (int, int) {
	if current-previous > 0 {
		counter++
	}

	return current, counter
}
