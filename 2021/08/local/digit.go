package local

import (
	"strings"
)

// Digit represent a number displayed on 7 segment indicator
type Digit struct {
	Segments []rune
	Decoded  int
}

func (d Digit) subtract(value Digit) []rune {
	temp := string(d.Segments)
	for _, search := range value.Segments {
		temp = strings.ReplaceAll(temp, string(search), "")
	}

	return []rune(temp)
}

func (d Digit) contains(value Digit) bool {
	if len(d.Segments)-len(value.Segments) == len(d.subtract(value)) {
		return true
	}

	return false
}
