package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/08/local"
)

func main() {
	input := local.Read("./input.txt")

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
	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i].Readings); j++ {
			length := len(input[i].Readings[j])
			// Digit `1` will lit up 2 segments
			// Digit `7` will lit up 3 segments
			// Digit `4` will lit up 4 segments
			// Digit `8` will lit up all 7 segments
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}

	result <- count
}

func solvePuzzleB(input local.Data, result chan int) {
	entries := append(local.Data(nil), input...)
	total := 0

	for i := 0; i < len(entries); i++ {
		decoder := make(map[int]local.Figure)
		segment := make(map[string]local.Figure)

		decoder[1] = entries[i].FindPatternsByLen(2)[0]
		decoder[4] = entries[i].FindPatternsByLen(4)[0]
		decoder[7] = entries[i].FindPatternsByLen(3)[0]
		decoder[8] = entries[i].FindPatternsByLen(7)[0]

		groupOf6 := entries[i].FindPatternsByLen(6)

		for _, digit := range groupOf6 {
			if digit.Contains(decoder[4]) {
				decoder[9] = digit
			}
		}

		for _, digit := range groupOf6 {
			// Group with length 6 contains 3 digits:
			// * `9` (which we identified previously, so we're excluding it here)
			// * `0` – it also fully contains `1`
			// * `6` that does not contain `1`
			if !digit.Equals(decoder[9]) {
				if digit.Contains(decoder[1]) {
					decoder[0] = digit
				} else {
					decoder[6] = digit
				}
			}
		}

		segment["A"] = decoder[7].Subtract(decoder[1])
		segment["C"] = decoder[8].Subtract(decoder[6])
		segment["D"] = decoder[8].Subtract(decoder[0])
		segment["E"] = decoder[8].Subtract(decoder[9])
		segment["F"] = decoder[1].Subtract(segment["C"])
		segment["G"] = decoder[9].Subtract(decoder[4]).Subtract(segment["A"])

		segment["B"] = decoder[8].Subtract(decoder[7]).Subtract(segment["G"]).Subtract(segment["D"]).Subtract(segment["E"])

		decoder[2] = decoder[8].Subtract(segment["B"]).Subtract(segment["F"])
		decoder[3] = decoder[8].Subtract(segment["B"]).Subtract(segment["E"])
		decoder[3] = decoder[8].Subtract(segment["C"]).Subtract(segment["E"])

		total += entries[i].GetValue(decoder)
	}

	result <- total
}
