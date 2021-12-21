package local

import "testing"

func TestDigitSubtract(t *testing.T) {
	// Arrange
	one := Figure{
		Segments: []rune{'c', 'f'},
		Decoded:  1,
	}
	seven := Figure{
		Segments: []rune{'a', 'c', 'f'},
		Decoded:  7,
	}

	expected := []rune{'a'}

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
	one := Figure{
		Segments: []rune{'c', 'f'},
		Decoded:  1,
	}
	four := Figure{
		Segments: []rune{'b', 'c', 'd', 'f'},
		Decoded:  4,
	}
	seven := Figure{
		Segments: []rune{'a', 'c', 'f'},
		Decoded:  7,
	}

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
