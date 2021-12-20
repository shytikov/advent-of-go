package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"math"
	"sync"
)

func main() {
	input := shared.ReadIntSliceFromLine("./input.txt")

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

func solvePuzzleA(input [][]int, result chan int) {
	positions := input[0]
	max := shared.MaxOf(positions)
	cost := make([]int, max+1)

	var wg sync.WaitGroup

	for i := 0; i <= max; i++ {
		wg.Add(1)

		go func(target int) {
			defer wg.Done()

			fuel := 0
			for _, position := range positions {
				fuel += int(math.Abs(float64(position - target)))
			}

			cost[target] = fuel
		}(i)
	}

	wg.Wait()

	frequency := make([]int, max+1)

	for _, position := range positions {
		frequency[position] += 1
	}

	result <- shared.MinOf(cost)
}

func solvePuzzleB(input [][]int, result chan int) {
	result <- 0
}
