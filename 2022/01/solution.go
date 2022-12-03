package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"sort"
)

func main() {
	input := shared.ReadSlice2D[int]("./input.txt", "\n\n", "\n")

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

func solvePart1(input [][]int, result chan int) {
	count := len(input)
	calories := make([]int, count)

	for i, entry := range input {
		value := 0
		for _, item := range entry {
			value += item
		}
		calories[i] = value
	}

	sort.Ints(calories)

	result <- calories[count-1]
}

func solvePart2(input [][]int, result chan int) {
	count := len(input)
	calories := make([]int, count)

	for i, entry := range input {
		value := 0
		for _, item := range entry {
			value += item
		}
		calories[i] = value
	}

	sort.Ints(calories)

	sum := 0

	for _, entry := range calories[count-3:] {
		sum += entry
	}

	result <- sum
}
