package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntGrid("./input.txt")

	if input != nil {
		answer1 := make(chan int)
		answer2 := make(chan int)

		go solvePart1(input[0], answer1)
		go solvePart2(input[0], answer2)

		fmt.Println(<-answer1)
		fmt.Println(<-answer2)
	} else {
		panic("failure when reading input data")
	}
}

func solvePart1(input []int, result chan int) {
	input = append(input, input[0])
	accumulator := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			accumulator += input[i]
		}
	}

	result <- accumulator
}

func solvePart2(input []int, result chan int) {
	input = append(input, input[0:len(input)/2]...)

	result <- 0
}
