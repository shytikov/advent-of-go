package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/08/local"
)

func main() {
	input := local.Read("./input.txt")

	if input != nil && len(input) > 0 {
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

func solvePuzzleA(input local.Data, result chan int) {
	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i].Readings); j++ {
			test := input[i].Readings[j].Decoded
			if test == 1 || test == 4 || test == 7 || test == 8 {
				count++
			}
		}
	}

	result <- count
}

func solvePuzzleB(input local.Data, result chan int) {
	result <- 0
}
