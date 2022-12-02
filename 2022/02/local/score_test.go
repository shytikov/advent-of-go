package local

import (
	"testing"
)

func TestScore(t *testing.T) {
	// Arrange
	input := [][]rune{
		{'B', 'X'},
		{'C', 'Y'},
		{'A', 'Z'},
		{'A', 'X'}, // Draw
		{'B', 'Y'}, // Draw
		{'C', 'Z'}, // Draw
		{'C', 'X'}, // Win
		{'A', 'Y'}, // Win
		{'B', 'Z'}, // Win
	}

	expected := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	// Act
	// Assert
	for i := 0; i < len(input); i++ {
		actual := Score(int(input[i][0])-64, int(input[i][1])-87)

		if actual != expected[i] {
			t.Errorf("Score of (%s, %s) was incorrect, got: %d, want: %d.",
				string(input[i][0]),
				string(input[i][1]),
				actual,
				expected[i])
			return
		}

	}
}
