package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/utils"
)

func main() {
	input := utils.ReadIntSlicesFromLines("./input.txt")

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
	//	_ := append([]int(nil), input[0]...)
	result <- 0
}
