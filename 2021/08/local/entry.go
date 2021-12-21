package local

import "math"

type Entry struct {
	Patterns []Figure
	Readings []Figure
	Segments map[rune][]rune
}

func (e Entry) FindPatternsByLen(value int) (result []Figure) {
	for _, digit := range e.Patterns {
		if len(digit) == value {
			result = append(result, digit)
		}
	}

	return
}

func (e Entry) GetValue(decoder map[int]Figure) (result int) {
	i := len(e.Readings)
	for _, reading := range e.Readings {
		i--
		for key, value := range decoder {
			if string(value) == string(reading) {
				result += key * int(math.Pow10(i))
			}
		}
	}

	return
}
