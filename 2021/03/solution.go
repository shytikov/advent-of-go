package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntGrid("./input.txt")

	if input != nil {
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

	result <- shared.LooseBinaryToInt(γ) * shared.LooseBinaryToInt(ε)
}

func solvePart2(input [][]int, result chan int) {
	gauge := map[string]chan []int{
		"O2":  make(chan []int),
		"CO2": make(chan []int),
	}

	O2 := func(count0, count1 int) bool { return count0 > count1 && count0 != count1 }
	CO2 := func(count0, count1 int) bool { return count0 < count1 || count0 == count1 }

	go calculateMetric(O2, input, gauge["O2"])
	go calculateMetric(CO2, input, gauge["CO2"])

	result <- shared.LooseBinaryToInt(<-gauge["O2"]) * shared.LooseBinaryToInt(<-gauge["CO2"])
}

func calculateMetric(detector func(count0 int, count1 int) bool, input [][]int, output chan []int) {
	notFiltered := input
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

		if detector(count0, count1) {
			notFiltered = filteredBy0
		} else {
			notFiltered = filteredBy1
		}

		if len(notFiltered) == 1 {
			output <- notFiltered[0]
			break
		}
	}
}
