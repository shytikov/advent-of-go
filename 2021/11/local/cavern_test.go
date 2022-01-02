package local

import (
	"fmt"
	"testing"
)

func TestCavernCreateGrid(t *testing.T) {
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
	actual := expected.CreateCavern().createGrid()

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

func TestCavernGetFlashesCount(t *testing.T) {
	// Arrange
	data := Data{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	lenX, lenY := len(data), len(data[0])
	chain := data.CreateCavern()
	expected := lenX * lenY

	// Act
	actual := chain.getFlashesCount()

	// Assert
	if actual != expected {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v", actual, expected)
		return
	}
}

func TestCavernAccumulateCharge(t *testing.T) {
	// Arrange
	data := Data{
		{1, 1, 1, 1, 1},
		{1, 9, 9, 9, 1},
		{1, 9, 1, 9, 1},
		{1, 9, 9, 9, 1},
		{1, 1, 1, 1, 1},
	}
	expected := []int{0, 9, 0}
	cavern := data.CreateCavern()

	// Act
	actual := make([]int, len(expected))
	actual[0] = cavern.getFlashesCount()

	for _, row := range cavern.createGrid() {
		fmt.Println(row)
	}

	fmt.Println()

	actual[1] = cavern.AccumulateCharge()

	for _, row := range cavern.createGrid() {
		fmt.Println(row)
	}

	fmt.Println()

	actual[2] = cavern.AccumulateCharge()

	for _, row := range cavern.createGrid() {
		fmt.Println(row)
	}

	fmt.Println()

	// Assert
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
			return
		}
	}
}
