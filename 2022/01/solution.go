package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"sort"
)

func main() {
	input := shared.ReadIntSlice2D("./input.txt", "\n\n", "\n")

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

func solvePuzzleA(input [][]int, result chan int) {
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

func solvePuzzleB(input [][]int, result chan int) {
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
