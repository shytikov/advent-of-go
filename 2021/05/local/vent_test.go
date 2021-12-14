package local

import "testing"

func TestIsOrthogonal(t *testing.T) {
	// Arrange
	vent := Vent{
		From: Point{
			599,
			531,
		},
		To: Point{
			599,
			32,
		},
	}

	expected := true

	// Act
	actual := vent.IsOrthogonal()

	// Assert
	if actual != expected {
		t.Errorf("Result was incorrect, got: %t, want: %t", actual, expected)
	}
}
