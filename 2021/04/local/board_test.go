package local

import (
	"strings"
	"testing"
)

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
	expected := 1142

	// Act
	board = board.Draw(17)
	actual := board.Sum

	// Assert
	if actual != expected {
		t.Errorf("Drawing present number should decrease the sum, got: %d, want: %d.", actual, expected)
	}
}
