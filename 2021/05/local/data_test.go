package local

import (
	"github.com/shytikov/advent-of-go/shared"
	"testing"
)

func TestParsePoint(t *testing.T) {
	// Arrange
	definition := "599,531"

	expected := shared.Point2D{
		X: 599,
		Y: 531,
	}

	// Act
	actual := parsePoint(definition)

	// Assert
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("Point value was incorrect, got: (%d, %d), want: (%d, %d).",
			actual.X,
			actual.Y,
			expected.X,
			expected.Y)
	}
}

func TestParseVent(t *testing.T) {
	// Arrange
	definition := "599,531 -> 599,32"

	expected := Vent{
		From: shared.Point2D{
			X: 599,
			Y: 531,
		},
		To: shared.Point2D{
			X: 599,
			Y: 32,
		},
	}

	// Act
	actual := parseVent(definition)

	// Assert
	if actual.From.X != expected.From.X || actual.From.Y != expected.From.Y {
		t.Errorf("`From` value was incorrect, got: (%d, %d), want: (%d, %d).",
			actual.From.X,
			actual.From.Y,
			expected.From.X,
			expected.From.Y)
	}

	if actual.To.X != expected.To.X || actual.To.Y != expected.To.Y {
		t.Errorf("`To` value was incorrect, got: (%d, %d), want: (%d, %d).",
			actual.To.X,
			actual.To.Y,
			expected.To.X,
			expected.To.Y)
	}
}

func TestParseData(t *testing.T) {
	// Arrange
	content :=
		`0,9 -> 5,9
		8,0 -> 0,8
		9,4 -> 3,4
		2,2 -> 2,1
		7,0 -> 7,4
		6,4 -> 2,0
		0,9 -> 2,9
		3,4 -> 1,4
		0,0 -> 8,8
		5,5 -> 8,2`

	expectedMin := shared.Point2D{}

	expectedMax := shared.Point2D{
		X: 9,
		Y: 9,
	}

	// Act
	actual := parseData(content)

	// Assert
	if actual.Min.X != expectedMin.X || actual.Min.Y != expectedMin.Y {
		t.Errorf("Min coordinate was wrong, got: (%d, %d), want: (%d, %d)", actual.Min.X, actual.Min.Y, expectedMin.X, expectedMin.Y)
	}

	if actual.Max.X != expectedMax.X || actual.Max.Y != expectedMax.Y {
		t.Errorf("Max coordinate was wrong, got: (%d, %d), want: (%d, %d)", actual.Max.X, actual.Max.Y, expectedMax.X, expectedMax.Y)
	}
}
