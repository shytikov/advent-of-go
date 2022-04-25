package local

import (
	"github.com/shytikov/advent-of-go/shared"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

type Data []Entry

func Read(filename string) Data {
	content, err := ioutil.ReadFile(filename)
	shared.ActOn(err)

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

func parsePart(definition string) (result []Figure) {
	chunks := strings.Split(definition, " ")

	result = make([]Figure, len(chunks))

	for i, digit := range chunks {
		result[i] = parseDigit(digit)
	}

	return
}

func parseDigit(definition string) (result Figure) {
	result = []rune(definition)

	// There is no particular need in sorting, but it will look more uniformly that way
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return
}
