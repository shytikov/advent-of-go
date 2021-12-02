package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadInstructions("./input.txt")

	if input != nil {
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

func solveFirstPuzzle(input []utils.Instruction, result chan int) {
	var path int
	var depth int

	for _, instruction := range input {
		switch instruction.Direction {
		case "forward":
			path += instruction.Distance
		case "down":
			depth += instruction.Distance
		case "up":
			if depth - instruction.Distance < 0 {
				depth = 0
			} else {
				depth -= instruction.Distance
			}
		}
	}

	result <- path*depth
}

func solveSecondPuzzle(input []utils.Instruction, result chan int) {
	result <- 2
}