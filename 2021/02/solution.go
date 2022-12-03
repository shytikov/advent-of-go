package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/02/local"
)

func main() {
	input := local.Read("./input.txt")

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

func solvePart1(input []local.Command, result chan int) {
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

func solvePart2(input []local.Command, result chan int) {
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
