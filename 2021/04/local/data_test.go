package local

import (
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
