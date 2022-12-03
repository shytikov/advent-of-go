package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/2021/12/local"
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

func solvePart1(input local.Data, result chan int) {
	result <- 0
}

func solvePart2(input local.Data, result chan int) {
	result <- 0
}
