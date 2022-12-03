package main

import (
	"fmt"

	"github.com/shytikov/advent-of-go/2021/04/local"
	"github.com/shytikov/advent-of-go/shared"
)

func main() {
	input := local.Read("./input.txt")

	if input.Draw != nil && input.Boards != nil {
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

func solvePart1(input local.Data, result chan int) {
	boards := append([]local.Board(nil), input.Boards...)

	for _, number := range input.Draw {
		for i, board := range boards {
			boards[i] = board.Draw(number)

			if boards[i].HasWon() {
				result <- boards[i].GetScore(number)

				return
			}
		}
	}

	result <- 0
}

func solvePart2(input local.Data, result chan int) {
	boards := append([]local.Board(nil), input.Boards...)

	count := len(boards)
	winners := make([]int, count)

	for _, number := range input.Draw {
		for i, board := range boards {
			boards[i] = board.Draw(number)

			if boards[i].HasWon() {
				winners[i] = 1
				if shared.SumOf(winners) == count {
					result <- boards[i].GetScore(number)
					return
				}
			}
		}
	}

	result <- 0
}
