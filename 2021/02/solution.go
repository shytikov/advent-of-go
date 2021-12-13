package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/02/local"
)

func main() {
	input := local.Read("./input.txt")

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

func solvePuzzleA(input []local.Command, result chan int) {
	stretch := 0
	depth := 0

	for _, instruction := range input {
		switch instruction.Name {
		case "forward":
			stretch += instruction.Argument
		case "down":
			depth += instruction.Argument
		case "up":
			depth -= instruction.Argument
		}
	}

	result <- stretch * depth
}

func solvePuzzleB(input []local.Command, result chan int) {
	stretch := 0
	depth := 0
	aim := 0

	for _, instruction := range input {
		switch instruction.Name {
		case "forward":
			stretch += instruction.Argument
			depth += aim * instruction.Argument
		case "down":
			aim += instruction.Argument
		case "up":
			aim -= instruction.Argument
		}
	}

	result <- stretch * depth
}
