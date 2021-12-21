package local

import (
	"strings"
)

// Figure represent a number displayed on 7 segment indicator
type Figure []rune

func (f Figure) Subtract(segments Figure) Figure {
	temp := string(f)
	for _, search := range segments {
		temp = strings.ReplaceAll(temp, string(search), "")
	}

	return []rune(temp)
}

func (f Figure) Contains(segments Figure) bool {
	if len(f)-len(segments) == len(f.Subtract(segments)) {
		return true
	}

	return false
}

func (f Figure) Equals(segments Figure) bool {
	if len(f.Subtract(segments)) == 0 {
		return true
	}

	return false
}
