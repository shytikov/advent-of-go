package shared

import "testing"

func TestParse1DPoint(t *testing.T) {
	// Arrange
	definition := "599"

	expected := Point{
		X: 599,
		Y: 0,
		Z: 0,
	}

	// Act
	actual := ParsePoint(definition)

	// Assert
	if actual.X != expected.X || actual.Y != expected.Y || actual.Z != expected.Z {
		t.Errorf("Point value was incorrect, got: (%d, %d, %d), want: (%d, %d, %d).",
			actual.X,
			actual.Y,
			actual.Z,
			expected.X,
			expected.Y,
			expected.Z,
		)
	}
}

func TestParse2DPoint(t *testing.T) {
	// Arrange
	definition := "599,531"

	expected := Point{
		X: 599,
		Y: 531,
		Z: 0,
	}

	// Act
	actual := ParsePoint(definition)

	// Assert
	if actual.X != expected.X || actual.Y != expected.Y || actual.Z != expected.Z {
		t.Errorf("Point value was incorrect, got: (%d, %d, %d), want: (%d, %d, %d).",
			actual.X,
			actual.Y,
			actual.Z,
			expected.X,
			expected.Y,
			expected.Z,
		)
	}
}

func TestParse3DPoint(t *testing.T) {
	// Arrange
	definition := "599,531,504"

	expected := Point{
		X: 599,
		Y: 531,
		Z: 504,
	}

	// Act
	actual := ParsePoint(definition)

	// Assert
	if actual.X != expected.X || actual.Y != expected.Y || actual.Z != expected.Z {
		t.Errorf("Point value was incorrect, got: (%d, %d, %d), want: (%d, %d, %d).",
			actual.X,
			actual.Y,
			actual.Z,
			expected.X,
			expected.Y,
			expected.Z,
		)
	}
}
