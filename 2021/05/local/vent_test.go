package local

import (
	"github.com/shytikov/advent-of-go/shared"
	"testing"
)

func TestIsOrthogonal(t *testing.T) {
	// Arrange
	vent := Vent{
		From: shared.Point2D{
			X: 599,
			Y: 531,
		},
		To: shared.Point2D{
			X: 599,
			Y: 32,
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
