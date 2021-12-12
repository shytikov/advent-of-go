package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseBoard(t *testing.T) {
	// Arrange
	rawString := `
	17 45 62 28 73
	39 12  0 52  5
	87 48 50 85 44
	66 57 78 94  3
	91 37 69 16  1`

	rawNumbers := [dimension][dimension]int{
		{17, 45, 62, 28, 73},
		{39, 12, 0, 52, 5},
		{87, 48, 50, 85, 44},
		{66, 57, 78, 94, 3},
		{91, 37, 69, 16, 1},
	}

	lines := strings.Split(rawString, "\n")

	// Act
	board := parseBoard(lines)

	// Assert
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			actual := board.Numbers[i][j]
			expected := rawNumbers[i][j]

			if actual != expected {
				t.Errorf("Number (%d, %d) was incorrect, got: %d, want: %d.", i, j, actual, expected)
				return
			}
		}
	}
}

func TestGetBoards(t *testing.T) {
	// Arrange
	content, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(content), "\n")
	expected := 100
	result := make(chan []Board, 1)

	// Act
	go getBoardsFrom(lines, result)

	boards := <-result
	total := len(boards)

	// Assert
	if total != expected {
		t.Errorf("Total count was incorrect, got: %d, want: %d.", total, expected)
		return
	}
}

func TestDrawNumber(t *testing.T) {
	// Arrange
	rawString := `
	17 45 62 28 73
	39 12  0 52  5
	87 48 50 85 44
	66 57 78 94  3
	91 37 69 16  1`

	lines := strings.Split(rawString, "\n")
	board := parseBoard(lines)

	// Act
	board = board.draw(17)

	// Assert
	fmt.Println(board)
}
