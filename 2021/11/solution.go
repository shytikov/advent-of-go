package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/11/local"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntGrid("./input.txt")

	if input != nil && len(input) > 0 {
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

func solvePart1(input local.Data, result chan int) {
	cavern := input.CreateArea()
	count := 0

	for i := 0; i < 100; i++ {
		count += cavern.AccumulateCharge()
	}

	result <- count
}

func solvePart2(input local.Data, result chan int) {
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
