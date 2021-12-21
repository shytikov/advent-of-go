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
			digit := input[i].Readings[j].Decoded
			if digit == 1 || digit == 4 || digit == 7 || digit == 8 {
				count++
			}
		}
	}

	result <- count
}

func solvePuzzleB(input local.Data, result chan int) {
	entries := append(local.Data(nil), input...)

	for i := 0; i < len(entries); i++ {
		a := entries[i].FindSegmentA()
		g := entries[i].FindSegmentG(a)

		fmt.Println(entries[i].FindPatternsByLen(2))
		fmt.Println(entries[i].FindPatternsByLen(3))
		fmt.Println(a)
		fmt.Println(g)

		break
	}

	result <- 0
}
