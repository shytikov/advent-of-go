package shared

import (
	"testing"
)

func TestReadIntSliceFromRuneSlice(t *testing.T) {
	// Arrange
	content :=
		`010101110000
		010011000110
		010101000011
		111100100001
		011100110101
		110001010101
		001111110101
		101100011100
		010111111011
		101010111101
		101000001110
		001000100001`

	expected := [][]int{
		{0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1},
		{1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
	}

	// Act
	actual := readIntSlicesFromRuneSlices(content)

	// Assert
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {

			if actual[i][j] != expected[i][j] {
				t.Errorf("Number (%d, %d) was incorrect, got: %d, want: %d.", i, j, actual[i][j], expected[i][j])
				return
			}
		}
	}
}

func TestReadIntSlicesFromLines(t *testing.T) {
	// Arrange
	content :=
		`1101,1,29,67,1102,0,1,65,1008,65,35,66
		1005,66,28,1,67,65,20,4,0,1001`

	expected := [][]int{
		{1101, 1, 29, 67, 1102, 0, 1, 65, 1008, 65, 35, 66},
		{1005, 66, 28, 1, 67, 65, 20, 4, 0, 1001},
	}

	// Act
	actual := readIntSlicesFromLines(content)

	// Assert
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {

			if actual[i][j] != expected[i][j] {
				t.Errorf("Number (%d, %d) was incorrect, got: %d, want: %d.", i, j, actual[i][j], expected[i][j])
				return
			}
		}
	}
}
