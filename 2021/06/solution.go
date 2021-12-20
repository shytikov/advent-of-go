package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadIntSlicesFromRuneSlices("./input.txt")

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

func solvePuzzleA(input [][]int, result chan int) {
	school := append([]int(nil), input[0]...)

	// This approach is not optimal, but its presents and actually ability to execute flawlessly on small amounts of data
	// is good for the contrast as it is not scale, and for second part another approach should be taken
	for i := 0; i < 80; i++ {
		count := len(school)

		for j := 0; j < count; j++ {
			switch school[j] {
			case 0:
				school = append(school, 8)
				school[j] = 6
			default:
				school[j] -= 1
			}
		}
	}

	result <- len(school)
}

func solvePuzzleB(input [][]int, result chan int) {
	states := 9
	school := make([]int, states)

	// This time instead of increasing slice length, hashmap is used that keeps count of all generations
	for _, fish := range input[0] {
		school[fish] += 1
	}

	for i := 0; i < 256; i++ {
		temp := school[0]

		for i := 1; i < states; i++ {
			school[i-1] = school[i]
		}

		school[6] += temp
		school[8] = temp
	}

	total := 0
	for _, count := range school {
		total += count
	}

	result <- total
}
