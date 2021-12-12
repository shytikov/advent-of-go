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
	boards := input.Boards

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
	// Skipping line with draw numbers
	lines = lines[1:]

	// Total total of boards could be found from number of lines left in input
	// by arranging them in groups of 6, 5x5 is demention of the board plus spacer line
	step := (dimension + 1)
	total := len(lines) / step

	boards := make([]Board, total)

	for i := 0; i < total; i++ {
		start := i * step
		finish := start + step

		boards[i] = parseBoard(lines[start:finish])
	}

	result <- boards
}

func parseBoard(lines []string) (result Board) {
	// Skipping spacer line
	lines = lines[1:]

	for i, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "  ", " ")
		numbers := strings.Split(line, " ")

		for j, number := range numbers {
			value, _ := strconv.Atoi(number)

			result.Numbers[i][j] = value
			result.Sum += value
		}
	}

	return
}

func (b Board) draw(number int) (result Board) {
	result = b

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if result.Numbers[i][j] == number {
				result.Rows[i] += 1
				result.Sum = result.Sum - number
			}
		}
	}

	return
}

func (b Board) hasWon() bool {
	for i := 0; i < dimension; i++ {
		if b.Rows[i] == dimension {
			return true
		}
	}

	return false
}

func (b Board) getScore(number int) int {
	return b.Sum * number
}
