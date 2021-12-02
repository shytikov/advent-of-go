package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadInstructionsFromLInes("./input.txt")

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

func solveFirstPuzzle(input []utils.Instruction, result chan int) {
	stretch := 0
	depth := 0

	for _, instruction := range input {
		switch instruction.Direction {
		case "forward":
			stretch += instruction.Distance
		case "down":
			depth += instruction.Distance
		case "up":
			depth -= instruction.Distance
		}
	}

	result <- stretch * depth
}

func solveSecondPuzzle(input []utils.Instruction, result chan int) {
	stretch := 0
	depth := 0
	aim := 0

	for _, instruction := range input {
		switch instruction.Direction {
		case "forward":
			stretch += instruction.Distance
			depth += aim * instruction.Distance
		case "down":
			aim += instruction.Distance
		case "up":
			aim -= instruction.Distance
		}
	}

	result <- stretch * depth
}
