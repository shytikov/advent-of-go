package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadIntSlicesFromLines("./input.txt")

	if input != nil {
		resultA := make(chan int)
		resultB := make(chan int)

		go solvePuzzleA(input, resultA)
		go solvePuzzleB(input, resultB)

		fmt.Println(<-resultA)
		fmt.Println(<-resultB)
	} else {
		fmt.Errorf("failure when reading input data")
	}
}

func solvePuzzleA(input [][]int, result chan int) {
	// Idea is to sum all values per digit position, and when compare with half-length of the input.
	// If sum exceeds it this means it was more `1` in this sum than `0`, and vice versa
	threshold := len(input) / 2
	length := len(input[0])

	// Creating a slice to accumulate sums of digits in each position
	accumulator := make([]int, length)
	for _, number := range input {
		for i, digit := range number {
			accumulator[i] = accumulator[i] + digit
		}
	}

	// Gamma – most common digit
	γ := make([]int, length)
	// Epsilon – least common digit
	ε := make([]int, length)
	for i, digit := range accumulator {
		if digit > threshold {
			γ[i] = 1
			ε[i] = 0
		} else {
			γ[i] = 0
			ε[i] = 1
		}
	}

	result <- utils.LooseBinaryToInt(γ) * utils.LooseBinaryToInt(ε)
}

func solvePuzzleB(input [][]int, result chan int) {
	notFiltered := input
	o2Level := 0
	co2Level := 0

	for i := 0; i < len(input[0]); i++ {
		filteredBy0 := make([][]int, 0)
		filteredBy1 := make([][]int, 0)

		count0 := 0
		count1 := 0

		for _, number := range notFiltered {
			if number[i] == 0 {
				count0++
				filteredBy0 = append(filteredBy0, number)
			} else {
				count1++
				filteredBy1 = append(filteredBy1, number)
			}
		}

		if count0 < count1 || count0 == count1 {
			notFiltered = filteredBy1
		} else {
			notFiltered = filteredBy0
		}

		if len(notFiltered) == 1 {
			o2Level = utils.LooseBinaryToInt(notFiltered[0])
			break
		}
	}

	notFiltered = input
	for i := 0; i < len(input[0]); i++ {
		filteredBy0 := make([][]int, 0)
		filteredBy1 := make([][]int, 0)

		count0 := 0
		count1 := 0

		for _, number := range notFiltered {
			if number[i] == 0 {
				count0++
				filteredBy0 = append(filteredBy0, number)
			} else {
				count1++
				filteredBy1 = append(filteredBy1, number)
			}
		}

		if count0 < count1 || count0 == count1 {
			notFiltered = filteredBy0
		} else {
			notFiltered = filteredBy1
		}

		if len(notFiltered) == 1 {
			co2Level = utils.LooseBinaryToInt(notFiltered[0])
			break
		}
	}

	result <- o2Level*co2Level
}
