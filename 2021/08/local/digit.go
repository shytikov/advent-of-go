package local

// Digit represent a number displayed on 7 segment indicator
type Digit struct {
	Original []rune
}

func (d Digit) getValue() int {
	switch len(d.Original) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	default:
		return -1
	}
}
