package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
	// for _, number := range input.Draw {
	// 	for _, board := range input.Boards {
	// 		board.draw(number)

	// 		if board.hasWon() {
	// 			result <- board.getScore(number)

	// 			return
	// 		}
	// 	}
	// }

	result <- 0
}

func solvePuzzleB(input Data, result chan int) {
	result <- 0
}

type Board struct {
	Numbers [dimension][dimension]int
	Rows    [dimension]int
	Sum     int
}

type Data struct {
	Draw   []int
	Boards []Board
}

func readBingo(filename string) Data {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return Data{}
	}

	lines := strings.Split(string(content), "\n")

	draw := make(chan []int)
	boards := make(chan []Board)

	go getDrawFrom(lines, draw)
	go getBoardsFrom(lines, boards)

	result := Data{
		Draw:   <-draw,
		Boards: <-boards,
	}

	return result
}

func getDrawFrom(lines []string, result chan []int) {
	raw := strings.Split(lines[0], ",")
	draw := make([]int, len(raw))

	for i, number := range raw {
		value, err := strconv.Atoi(number)

		if err != nil {
			continue
		}
		draw[i] = value
	}

	result <- draw
}

func getBoardsFrom(lines []string, result chan []Board) {

	result <- nil
}

func parseBoard(lines []string) (result Board) {

	return
}

// func (b Board) draw(number int) {
// 	for i := 0; i < dimension; i++ {
// 		for j := 0; j < dimension; j++ {
// 			if b.Numbers[i][j] == number {
// 				b.Sum[i][j] = 1
// 				b.Rows[i] += 1
// 			}
// 		}
// 	}
// }

// func (b Board) hasWon() (result bool) {
// 	for i := 0; i < dimension; i++ {
// 		if utils.SumInts(b.Sum[i]) == dimension {
// 			result = true
// 			break
// 		}
// 	}

// 	return
// }

// func (b Board) getScore(number int) (result int) {
// 	for i := 0; i < dimension; i++ {
// 		result += utils.SumInts(b.Rows[i])
// 	}

// 	return result * number
// }
