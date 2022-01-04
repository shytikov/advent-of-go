package local

import (
	"testing"
)

func TestDataCreateArea(t *testing.T) {
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

	expectedLength := lenX * lenY

	// Act
	actual := expected.CreateArea()

	// Assert
	if len(actual) != expectedLength {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v",
			len(actual),
			expectedLength)
		return
	}

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			octopus := actual[expected.getOctopusIndex(i, j)]
			if octopus.Value.Z != expected[i][j] {
				t.Errorf("Result was incorrect for index (%v,%v), got: %v, want: %v",
					i,
					j,
					octopus.Value.Z,
					expected[i][j])
				return
			}
		}
	}
}

func TestDataGetNeighbourCount(t *testing.T) {
	// Arrange
	expected := Data{
		{3, 5, 5, 5, 5, 5, 5, 5, 5, 3},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{5, 8, 8, 8, 8, 8, 8, 8, 8, 5},
		{3, 5, 5, 5, 5, 5, 5, 5, 5, 3},
	}
	lenX, lenY := len(expected), len(expected[0])

	// Act
	// Assert
	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			length := expected.getNeighbourCount(i, j)
			if length != expected[i][j] {
				t.Errorf("Result was incorrect for index (%v,%v), got: %v, want: %v", i, j, length, expected[i][j])
				return
			}
		}
	}
}

func TestDataGetNeighbourIndexes(t *testing.T) {
	// Arrange
	data := Data{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{70, 71, 72, 73, 74, 75, 76, 77, 78, 79},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
	}

	expected := map[int][]int{
		0:  []int{1, 10, 11},
		1:  []int{0, 2, 10, 11, 12},
		11: []int{0, 1, 2, 10, 12, 20, 21, 22},
		88: []int{77, 78, 79, 87, 89, 97, 98, 99},
	}

	// Act
	actual := map[int][]int{
		0:  data.getNeighbourIndexes(0, 0),
		1:  data.getNeighbourIndexes(0, 1),
		11: data.getNeighbourIndexes(1, 1),
		88: data.getNeighbourIndexes(8, 8),
	}

	// Assert
	for k, v := range actual {
		if len(v) != len(expected[k]) {
			t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v", len(v), len(expected[k]))
			return
		}

		for i, item := range v {
			if item != expected[k][i] {
				t.Errorf("Result was incorrect, neighbours don't match for index %v, got: %v, want: %v", k, v, expected[k])
				return
			}
		}
	}
}

func TestDataGetOctopusIndex(t *testing.T) {
	// Arrange
	expected := Data{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{70, 71, 72, 73, 74, 75, 76, 77, 78, 79},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
	}

	lenX, lenY := len(expected), len(expected[0])

	// Act
	actual := make([][]int, lenX)

	for i := 0; i < lenX; i++ {
		actual[i] = make([]int, lenY)
		for j := 0; j < lenY; j++ {
			actual[i][j] = expected.getOctopusIndex(i, j)
		}
	}

	// Assert
	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			if actual[i][j] != expected[i][j] {
				t.Errorf("Result was incorrect for index (%v,%v), got: %v, want: %v", i, j, actual[i][j], expected[i][j])
				return
			}
		}
	}
}
