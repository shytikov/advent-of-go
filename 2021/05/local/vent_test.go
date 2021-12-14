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

func TestVentIsHorizontal(t *testing.T) {
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

	expectedHorizontal := true
	expectedVertical := false

	// Act
	actualHorizontal := vent.IsHorizontal()
	actualVertical := vent.IsVertical()

	// Assert
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Result was incorrect, got: %t, want: %t", actualHorizontal, expectedHorizontal)
	}

	if actualVertical != expectedVertical {
		t.Errorf("Result was incorrect, got: %t, want: %t", actualVertical, expectedVertical)
	}
}

func TestVentIsVertical(t *testing.T) {
	// Arrange
	vent := Vent{
		From: Point{
			771,
			406,
		},
		To: Point{
			120,
			406,
		},
	}

	expectedHorizontal := false
	expectedVertical := true

	// Act
	actualHorizontal := vent.IsHorizontal()
	actualVertical := vent.IsVertical()

	// Assert
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Result was incorrect, got: %t, want: %t", actualHorizontal, expectedHorizontal)
	}

	if actualVertical != expectedVertical {
		t.Errorf("Result was incorrect, got: %t, want: %t", actualVertical, expectedVertical)
	}
}
