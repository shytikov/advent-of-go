package local

import (
	"io/ioutil"
	"strings"
	"sync"
)

type Data []Entry

func Read(filename string) Data {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil
	}

	return parseData(string(content))
}

func parseData(content string) (result Data) {
	lines := strings.Split(content, "\n")
	count := len(lines)

	result = make(Data, count)

	var wg sync.WaitGroup

	for i, line := range lines {
		wg.Add(1)

		go func(index int, definition string) {
			defer wg.Done()
			chunks := strings.Split(definition, " | ")

			result[index] = Entry{
				Patterns: parsePart(chunks[0]),
				Readings: parsePart(chunks[1]),
			}

		}(i, line)
	}

	wg.Wait()

	return
}

func parsePart(definition string) (result []Digit) {
	chunks := strings.Split(definition, " ")

	result = make([]Digit, len(chunks))

	for i, digit := range chunks {
		result[i] = parseDigit(digit)
	}

	return
}

func parseDigit(definition string) (result Digit) {
	result.Original = definition

	switch len(definition) {
	case 2:
		result.Decoded = 1
	case 3:
		result.Decoded = 7
	case 4:
		result.Decoded = 4
	case 7:
		result.Decoded = 8
	default:
		result.Decoded = -1
	}

	return
}
