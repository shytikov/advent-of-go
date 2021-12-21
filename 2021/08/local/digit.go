package local

// Digit represent a number displayed on 7 segment indicator
type Digit struct {
	Segments []rune
	Decoded  int
}

func (d Digit) subtract(value Digit) (result []rune) {
	for _, search := range value.Segments {
		for _, current := range d.Segments {
			if current != search {
				result = append(result, current)
			}
		}
	}

	return
}
