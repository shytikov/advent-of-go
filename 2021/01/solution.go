package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
	"math"
)

func main() {
	input := utils.ReadIntFromLines("./input.txt")

	first := func() int {
		// Resetting counter
		var counter int

		// Previous depth value – intentionally set to big number
		previous := math.MaxInt16
		for _, current := range input {
			previous, counter = decide(current, previous, counter)
		}

		return counter
	}

	second := func() int {
		// Resetting counter
		var counter int

		// Previous depth value – intentionally set to big number
		previous := math.MaxInt16
		for i := 0; i < len(input)-2; i++ {
			previous, counter = decide(utils.SumInt(input[i:i+3]), previous, counter)
		}

		return counter
	}

	if input != nil {
		fmt.Println(first())
		fmt.Println(second())
	}
}

// FUnction will decide which value should be used for counter
func decide(current, previous, counter int) (int, int) {
	if current-previous > 0 {
		counter++
	}

	return current, counter
}
