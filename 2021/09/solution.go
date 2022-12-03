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
	points := input.DetectLowestPoints()
	sum := 0

	for _, coordinate := range points {
		sum += coordinate.Z
	}

	// Calculating risk. Risk = height of a point + one.
	// But we can sum all heights and add as many ones as length of the height slice is
	result <- sum + len(points)
}

func solvePart2(input local.Data, result chan int) {
	sizes := input.DetectBasinSizes(input.DetectLowestPoints())

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result <- shared.ProductOf(sizes[:3])
}
