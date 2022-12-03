package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"math"
	"sync"
)

func main() {
	input := shared.ReadSlice2D[int]("./input.txt", "\n", ",")

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

func solvePart1(input [][]int, result chan int) {
	// Constant cost function
	costFunction := func(distance int) int {
		return int(math.Abs(float64(distance)))
	}

	cost := calculateCost(input[0], costFunction)

	result <- shared.MinOf(cost)
}

func solvePart2(input [][]int, result chan int) {
	// Linear cost function
	costFunction := func(distance int) (result int) {
		cost := 0
		for i := 0; i <= int(math.Abs(float64(distance))); i++ {
			cost += i
		}
		result = cost

		return
	}

	cost := calculateCost(input[0], costFunction)

	result <- shared.MinOf(cost)
}

func calculateCost(positions []int, costFunction func(int) int) (result []int) {
	max := shared.MaxOf(positions)
	result = make([]int, max+1)

	var wg sync.WaitGroup

	for i := 0; i <= max; i++ {
		wg.Add(1)

		go func(target int) {
			defer wg.Done()

			fuel := 0
			for _, position := range positions {
				fuel += costFunction(position - target)
			}

			result[target] = fuel
		}(i)
	}

	wg.Wait()

	return
}
