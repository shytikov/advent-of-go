package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"math"
)

func main() {
	input := shared.ReadIntSlice1D("./input.txt", "\n")

	if input != nil {
		answer1 := make(chan int)
		answer2 := make(chan int)

		go solvePart1(input, answer1)
		go solvePart2(input, answer2)

		fmt.Println(<-answer1)
		fmt.Println(<-answer2)
	} else {
		panic("failure when reading input data")
	}
}

func solvePart1(input []int, result chan int) {
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

func solvePart2(input []int, result chan int) {
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
