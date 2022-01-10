package local

import (
	"testing"

	"github.com/shytikov/advent-of-go/shared"
)

func TestDetectLowestPoints(t *testing.T) {
	// Arrange
	data := Data{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	expected := []shared.Point{
		{X: 0, Y: 1, Z: 1},
		{X: 0, Y: 9, Z: 0},
		{X: 2, Y: 2, Z: 5},
		{X: 4, Y: 6, Z: 5},
	}

	// Act
	actual := data.DetectLowestPoints()

	// Assert
	if len(actual) != len(expected) {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v", len(actual), len(expected))
		return
	}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
			return
		}
	}
}

func TestDetectBasinSizes(t *testing.T) {
	// Arrange
	data := Data{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	coordinates := data.DetectLowestPoints()

	expected := []int{3, 9, 14, 9}

	// Act
	actual := data.DetectBasinSizes(coordinates)

	// Assert
	if len(actual) != len(expected) {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v", len(actual), len(expected))
		return
	}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
			return
		}
	}
}
