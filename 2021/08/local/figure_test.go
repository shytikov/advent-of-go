package local

import "testing"

func TestDigitSubtract(t *testing.T) {
	// Arrange
	one := Figure{'c', 'f'}
	seven := Figure{'a', 'c', 'f'}

	expected := Figure{'a'}

	// Act
	actual := seven.Subtract(one)

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

func TestDigitContains(t *testing.T) {
	// Arrange
	one := Figure{'c', 'f'}
	four := Figure{'b', 'c', 'd', 'f'}
	seven := Figure{'a', 'c', 'f'}

	expected := []bool{
		true,
		false,
	}

	// Act
	actual := []bool{
		seven.Contains(one),
		four.Contains(seven),
	}

	// Assert
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
			return
		}
	}
}


func TestDigitEquals(t *testing.T) {
	// Arrange
	one := Figure{'c', 'f'}
	four := Figure{'b', 'c', 'd', 'f'}

	expected := []bool{
		true,
		false,
	}

	// Act
	actual := []bool{
		one.Contains(Figure{'c', 'f'}),
		four.Equals(one),
	}

	// Assert
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
			return
		}
	}
}
