package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
	"math"
)

func main() {
	input := utils.ReadNumbers("./input.txt")

	if input != nil {
		firstResult := make(chan int)
		secondResult := make(chan int)

		go solveFirstPuzzle(input, firstResult)
		go solveSecondPuzzle(input, secondResult)

		fmt.Println(<-firstResult)
		fmt.Println(<-secondResult)
	} else {
		fmt.Errorf("failure when reading input data")
	}
}

func solveFirstPuzzle(input []int, result chan int) {
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

func solveSecondPuzzle(input []int, result chan int) {
	counter := 0
	previous := math.MaxInt16
	// It does not matter would we use len(input) or len(input)-2 as missing elements will be replaced with 0
	// and calculation will be still correct, but in this case we will run two unnecessary loop cycles
	for i := 0; i < len(input)-2; i++ {
		current := utils.SumNumbers(input[i : i+3])
		if current-previous > 0 {
			counter++
		}
		previous = current
	}

	result <- counter
}
