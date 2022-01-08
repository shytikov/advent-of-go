package local

import (
	"testing"
)

func TestDataRead(t *testing.T) {
	// Arrange
	content := `
		start-A
		start-b
		A-c
		A-b
		b-d
		A-end
		b-end`

	expected := 2

	// Act
	actual := read(content)

	// Assert
	if len(actual.Links) != expected {
		t.Errorf("Result was incorrect, lengths don't match, got: %v, want: %v",
			len(actual.Links),
			expected)
		return
	}
}
