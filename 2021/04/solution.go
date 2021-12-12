package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/utils"
)

const dimension = 5

func main() {
	input := readBingo("./input.txt")

	if input.Draw != nil && input.Boards != nil {
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

func solvePuzzleA(input Data, result chan int) {
	boards := append([]Board(nil), input.Boards...)

	for _, number := range input.Draw {
		for i, board := range boards {
			boards[i] = board.draw(number)

			if boards[i].hasWon() {
				result <- boards[i].getScore(number)

				return
			}
		}
	}

	result <- 0
}

func solvePuzzleB(input Data, result chan int) {
	boards := append([]Board(nil), input.Boards...)

	count := len(boards)
	winners := make([]int, count)

	for _, number := range input.Draw {
		for i, board := range boards {
			boards[i] = board.draw(number)

			if boards[i].hasWon() {
				winners[i] = 1
				if utils.SumInts(winners) == count {
					result <- boards[i].getScore(number)
					return
				}
			}
		}
	}

	result <- 0
}
