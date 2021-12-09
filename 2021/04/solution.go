package main

import (
	"fmt"
	"github.com/shytikov/advent-of-go/utils"
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
		fmt.Errorf("failure when reading input data")
	}
}

func solvePuzzleA(input Data, result chan int) {
	for _, number := range input.Draw {
		for _, board := range input.Boards {
			board.draw(number)

			if board.hasWon() {
				result <- board.getScore(number)

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
	Initial  [][]int
	Unmarked [][]int
	Marked   [][]int
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

	//fmt.Println(result)

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
	// A board is of size 5x5, but with one separation line in-front it would be 6 lines.
	// If we skipped first file with draw result rest of the input should be multiple of 6
	step := dimension + 1
	size := len(lines[1:]) / step
	boards := make([]Board, size)

	// Iterating over boards
	for i := 0; i < size; i++ {
		start := i*step + 1
		end := start + step
		board := Board{
			Initial:  make([][]int, dimension),
			Unmarked: make([][]int, dimension),
			Marked:   make([][]int, dimension),
		}

		j := 0

		// Iterating over lines within one board
		for k := start; k < end; k++ {
			// Skipping separator line
			if lines[k] != "" {
				row := make([]int, 5)
				// There is a padding in input data, it should be trimmed
				line := strings.Trim(lines[k], " ")
				line = strings.Replace(line, "  ", " ", 10)

				// Iterating over numbers in the line
				for l, number := range strings.Split(line, " ") {
					value, err := strconv.Atoi(number)

					if err != nil {
						continue
					}

					row[l] = value
				}

				board.Initial[j] = row
				board.Unmarked[j] = row
				board.Marked[j] = make([]int, dimension)
				j++
			}
		}

		boards[i] = board
	}

	result <- boards
}

func (b Board) draw(number int) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if b.Initial[i][j] == number {
				b.Marked[i][j] = 1
				b.Unmarked[i][j] = 0
			}
		}
	}
}

func (b Board) hasWon() (result bool) {
	for i := 0; i < dimension; i++ {
		if utils.SumInts(b.Marked[i]) == dimension {
			result = true
			break
		}
	}

	return
}

func (b Board) getScore(number int) (result int) {
	for i := 0; i < dimension; i++ {
		result += utils.SumInts(b.Unmarked[i])
	}

	return result * number
}
