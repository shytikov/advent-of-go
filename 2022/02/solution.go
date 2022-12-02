package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/2022/02/local"
	"github.com/shytikov/advent-of-go/shared"
	"sync"
)

func main() {
	input := shared.ReadSlice2D[rune]("./input.txt", "\n", " ")

	if input != nil {
		resultA := make(chan int)
		//resultB := make(chan int)

		go solvePuzzleA(input, resultA)
		//go solvePuzzleB(input, resultB)
		//
		fmt.Println(<-resultA)
		//fmt.Println(<-resultB)
	} else {
		panic("failure when reading input data")
	}
}

func solvePuzzleA(input [][]rune, result chan int) {
	total := make(chan int, len(input))

	var wg sync.WaitGroup

	for _, entry := range input {
		wg.Add(1)

		go func(draws []rune) {
			defer wg.Done()

			theirs := int(draws[0]) - 64
			ours := int(draws[1]) - 87

			total <- local.Score(theirs, ours)
		}(entry)
	}

	wg.Wait()
	close(total)

	result <- shared.GetSumFromChannel(total)
}

func solvePuzzleB(input [][]int, result chan int) {
	count := len(input)

	result <- count
}
