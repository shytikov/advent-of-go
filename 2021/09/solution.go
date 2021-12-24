package main

import (
	"fmt"
	"sort"

	"github.com/shytikov/advent-of-go/2021/09/local"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := local.Data(shared.ReadIntSlicesFromRuneSlices("./input.txt"))

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
	_, heights := input.DetectLowestPoints()

	// Calculating risk. Risk = height of a point + one.
	// But we can sum all heights and add as many ones as length of the height slice is
	result <- shared.SumOf(heights) + len(heights)
}

func solvePuzzleB(input local.Data, result chan int) {
	coordinates, _ := input.DetectLowestPoints()

	sizes := input.DetectBasinSizes(coordinates)

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	//fmt.Println(sizes[:3])

	result <- shared.ProductOf(sizes[:3])
}
