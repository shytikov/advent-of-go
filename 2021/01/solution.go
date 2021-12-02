package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
	"math"
	"strconv"
)

func main() {
	input := utils.ReadLines("./input.txt")
	fmt.Println(first(input))
}

func first(input []string) int {
	// Counter for increased depth input
	counter := 0

	// Previous depth value – intentionally set to big number
	previous := math.MaxInt16
	for _, measurement := range input {
		current, err := strconv.Atoi(measurement)

		if err != nil {
			panic(fmt.Errorf("cannot convert measurement to number"))
		}

		if current-previous > 0 {
			counter++
		}

		previous = current
	}

	return counter
}