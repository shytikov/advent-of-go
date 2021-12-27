package main

import (
	"fmt"
	"sort"

	"github.com/shytikov/advent-of-go/2021/09/local"
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
	sum, count := 0, 0

	for _, coordinate := range input.DetectLowestPoints() {
		sum += coordinate.Value
		count++
	}

	// Calculating risk. Risk = height of a point + one.
	// But we can sum all heights and add as many ones as length of the height slice is
	result <- sum + count
}

func solvePuzzleB(input local.Data, result chan int) {
	sizes := input.DetectBasinSizes(input.DetectLowestPoints())

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result <- shared.ProductOf(sizes[:3])
}
