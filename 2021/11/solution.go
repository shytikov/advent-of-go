package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/11/local"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntGrid("./input.txt")

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
	cavern := input.CreateArea()
	count := 0

	for i := 0; i < 100; i++ {
		count += cavern.AccumulateCharge()
	}

	result <- count
}

func solvePuzzleB(input local.Data, result chan int) {
	cavern := input.CreateArea()

	size := len(cavern)
	step := 0

	for {
		step++

		if cavern.AccumulateCharge() == size {
			break
		}
	}

	result <- step
}
