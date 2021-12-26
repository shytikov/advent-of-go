package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := shared.ReadRuneSlicesFromLines("./input.txt")

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

func solvePuzzleA(input [][]rune, result chan int) {
	corrupted, _ := split(input)
	cost := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	total := 0

	for _, current := range corrupted {
		total += cost[current]
	}

	result <- total
}

func solvePuzzleB(input [][]rune, result chan int) {
	_, incomplete := split(input)
	cost := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scores := make([]int, len(incomplete))

	for i, current := range incomplete {
		total := 0
		for _, item := range current {
			total *= 5
			total += cost[item]
		}
		scores[i] = total
	}

	sort.Ints(scores)

	result <- scores[(len(scores)-1)/2]
}

func split(input [][]rune) (corrupted []rune, incomplete []string) {
	// Collection of lookup objects: for opening, closing braces
	opening := "([{<"
	closing := ")]}>"

	for _, readings := range input {
		// Stack to keep track of braces that should come next
		stack := ""
		for _, current := range readings {
			index := strings.Index(opening, string(current))

			if index >= 0 {
				// It was an opening brace (index is positive),
				// adding respective closing one to the stack
				stack = string(closing[index]) + stack
			} else if rune(stack[0]) == current {
				// It was a closing brace that we expected,
				// as it matches with the first item in stack
				stack = stack[1:]
			} else if rune(stack[0]) != current {
				corrupted = append(corrupted, current)
				stack = ""
				break
			}
		}

		if len(stack) != 0 {
			incomplete = append(incomplete, stack)
		}
	}

	return
}
