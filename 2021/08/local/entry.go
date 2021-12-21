package local

type Entry struct {
	Patterns []Digit
	Readings []Digit
	Segments map[rune][]rune
}

func (e Entry) FindPatternsByLen(value int) (result []Digit) {
	for _, digit := range e.Patterns {
		if len(digit.Segments) == value {
			result = append(result, digit)
		}
	}

	return
}

func (e Entry) FindSegmentA() rune {
	// `Patterns` always contain all the digits, and only once,
	// so we might be sure that `1` and `7` will be found
	one := e.FindPatternsByLen(2)[0]
	seven := e.FindPatternsByLen(3)[0]

	return seven.subtract(one)[0]
}

func (e Entry) FindSegmentG(a rune) rune {
	// `Patterns` always contain all the digits, and only once
	// so we might be sure that `1` and `7` will be found
	one := e.FindPatternsByLen(2)[0]
	seven := e.FindPatternsByLen(3)[0]

	return seven.subtract(one)[0]
}
