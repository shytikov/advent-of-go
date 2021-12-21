package local

type Entry struct {
	Patterns []Digit
	Readings []Digit
	Segments map[int][]rune
}
