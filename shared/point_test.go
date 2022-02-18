package shared

import "testing"

func TestParsePoint(t *testing.T) {
	// Arrange
	definition := "599,531"

	expected := Point{
		X: 599,
		Y: 531,
	}

	// Act
	actual := ParsePoint(definition)

	// Assert
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("Point value was incorrect, got: (%d, %d), want: (%d, %d).",
			actual.X,
			actual.Y,
			expected.X,
			expected.Y)
	}
}

