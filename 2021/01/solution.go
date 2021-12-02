package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
	"math"
)

func main() {
	input := utils.ReadIntFromLines("./input.txt")

	if input != nil {
		fmt.Println(first(input))
		fmt.Println(second(input))
	}
}

func first(input []int) int {
	// Counter for increased depth samples count
	counter := 0

	// Previous depth value – intentionally set to big number
	previous := math.MaxInt16
	for _, current := range input {
		if current-previous > 0 {
			counter++
		}
		previous = current
	}

	return counter
}

func second(input []int) int {
	// Counter for increased depth samples count
	counter := 0

	// Previous depth value – intentionally set to big number
	previous := math.MaxInt16
	for i := 0; i < len(input)-2; i++ {
		current := utils.SumInt(input[i : i+3])

		if current-previous > 0 {
			counter++
		}
		previous = current
	}

	return counter
}
