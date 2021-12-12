package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

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

	return Data{
		Draw:   <-draw,
		Boards: <-boards,
	}
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
