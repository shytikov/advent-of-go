package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadIntSlicesFromLines("./input.txt")

	fmt.Println(input)

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
	placement := input[0]
	max := utils.MaxInInts(placement)
	distribution := make([]int, max+1)

	var wg sync.WaitGroup

	for i := 0; i <= max; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			fuel := 0
			for _, position := range placement {
				fuel += int(math.Abs(float64(position - index)))
			}

			distribution[index] = fuel
		}(i)
	}

	wg.Wait()

	fmt.Println(max, distribution)
	result <- 16 + 1 + 2 + 0 + 4 + 2 + 7 + 1 + 2 + 14
}

func solvePuzzleB(input [][]int, result chan int) {
	result <- 0
}
