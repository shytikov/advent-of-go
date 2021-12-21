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

	for i := 0; i < len(entries); i++ {
		// zero := local.Figure{}
		one := entries[i].FindPatternsByLen(2)[0]
		//four := entries[i].FindPatternsByLen(4)[0]
		seven := entries[i].FindPatternsByLen(3)[0]

		a := seven.Subtract(one)[0]

		//nine := local.Figure{}
		//
		//for _, digit := range entries[i].FindPatternsByLen(6) {
		//	if digit.Contains(four) {
		//		nine = digit
		//	}
		//}

		fmt.Println(a)
		//fmt.Println(g)

		break
	}

	result <- 0
}
