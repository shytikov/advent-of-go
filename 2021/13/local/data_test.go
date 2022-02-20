package local

import (
	"testing"

	"github.com/shytikov/advent-of-go/shared"
)

func TestDataRead(t *testing.T) {
	// Arrange
	content := `
		6,10
		0,14
		9,10
		0,3
		10,4
		4,11
		6,0
		6,12
		4,1
		0,13
		10,12
		3,4
		3,0
		8,4
		1,10
		2,14
		8,10
		9,0
		
		fold along y=7
		fold along x=5`

	expected := Data{
		Points: []shared.Point{
			{6, 10, 0},
			{0, 14, 0},
			{9, 10, 0},
			{0, 3, 0},
			{10, 4, 0},
			{4, 11, 0},
			{6, 0, 0},
			{6, 12, 0},
			{4, 1, 0},
			{0, 13, 0},
			{10, 12, 0},
			{3, 4, 0},
			{3, 0, 0},
			{8, 4, 0},
			{1, 10, 0},
			{2, 14, 0},
			{8, 10, 0},
			{9, 0, 0},
		},
		Folds: []shared.Point{
			{0, 7, 0},
			{5, 0, 0},
		},
	}

	// Act
	actual := parseData(content)

	// Assert
	if len(actual.Points) != len(expected.Points) {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v",
			len(actual.Points),
			len(expected.Points))
		return
	}

	for i := 0; i < len(actual.Points); i++ {
		if actual.Points[i].String() != expected.Points[i].String() {
			t.Errorf("Result was incorrect, items don't match, got: %v, want: %v",
				actual.Points[i],
				expected.Points[i])
		}
	}

	if len(actual.Folds) != len(expected.Folds) {
		t.Errorf("Point information was incorrect, lengths don't match, got: %v, want: %v",
			len(actual.Folds),
			len(expected.Folds))
		return
	}

	for i := 0; i < len(actual.Folds); i++ {
		if actual.Folds[i].String() != expected.Folds[i].String() {
			t.Errorf("Fold information was incorrect, items don't match, got: %v, want: %v",
				actual.Folds[i],
				expected.Folds[i])
		}
	}

}
