package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadNumbersFromString("./input.txt")

	if input != nil {
		firstResult := make(chan int)
		secondResult := make(chan int)

		input := append(input, input[0])

		go solveFirstPuzzle(input, firstResult)
		go solveSecondPuzzle(input, secondResult)

		fmt.Println(<-firstResult)
		fmt.Println(<-secondResult)
	} else {
		fmt.Errorf("failure when reading input data")
	}
}

func solveFirstPuzzle(input []int, result chan int) {
	accumulator := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			accumulator += input[i]
		}
	}

	result <- accumulator
}

func solveSecondPuzzle(input []int, result chan int) {
	result <- 0
}
