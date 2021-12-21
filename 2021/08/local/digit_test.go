package local

import "testing"

func TestDigitSubtract(t *testing.T) {
	// Arrange
	one := Digit{
		Segments: []rune{'c', 'f'},
		Decoded: 1,
	}
	seven := Digit{
		Segments: []rune{'a', 'c','f'},
		Decoded: 7,
	}

	expected := []rune {'a'}

	// Act
	actual := seven.subtract(one)

	// Assert
	if len(actual) != len(expected) {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v", len(actual), len(expected))
	}

	for i := 0; i< len(expected); i++  {
		if actual[i] != expected[i] {
			t.Errorf("Result was incorrect, got: %v, want: %v", actual[i], expected[i])
		}
	}
}
