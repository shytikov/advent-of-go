package local

import (
	"testing"
)

func TestChainCreateGrid(t *testing.T) {
	// Arrange
	expected := Data{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
	lenX, lenY := len(expected), len(expected[0])

	// Act
	actual := expected.CreateChain().createGrid()

	// Assert
	if len(actual) != len(expected) {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v", len(actual), len(expected))
		return
	}

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Result was incorrect, got: %v, want: %v", actual[i][j], expected[i][j])
				return
			}

		}
	}
}
