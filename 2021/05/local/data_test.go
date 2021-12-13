package local

import (
	"testing"
)

func TestParsePoint(t *testing.T) {
	// Arrange
	definition := "599,531"

	expected := Point{
		599,
		531,
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
		From: Point{
			599,
			531,
		},
		To: Point{
			599,
			32,
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
