package shared

import (
	"os"
	"strconv"
	"strings"
)

func ReadIntSlice1D(filename, sep string) (result []int) {
	content, err := os.ReadFile(filename)
	ActOn(err)

	for _, line := range strings.Split(string(content), sep) {
		number, err := strconv.Atoi(line)

		if err != nil {
			ActOn(err)
		}

		result = append(result, number)
	}

	if len(result) == 0 {
		return nil
	}

	return
}

// ReadIntGrid will read file and will treat each line as slice of runes
// Each rune then will be converted to integer
func ReadIntGrid(filename string) (result [][]int) {
	content, err := os.ReadFile(filename)
	ActOn(err)

	result = readIntGrid(string(content))

	if result == nil || len(result) == 0 {
		return nil
	}

	return
}

func readIntGrid(content string) (result [][]int) {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		chars := []rune(line)
		series := make([]int, len(chars))
		for i, char := range chars {
			number, err := strconv.Atoi(string([]rune{char}))
			ActOn(err)

			series[i] = number
		}

		result = append(result, series)
	}

	return
}

func ReadRuneSlicesFromLines(filename string) (result [][]rune) {
	content, err := os.ReadFile(filename)
	ActOn(err)

	result = readRuneSlicesFromLines(string(content))

	if result == nil || len(result) == 0 {
		return nil
	}

	return
}

func readRuneSlicesFromLines(content string) (result [][]rune) {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		result = append(result, []rune(line))
	}

	return
}

func ReadIntSlice2D(filename, sep1, sep2 string) (result [][]int) {
	content, err := os.ReadFile(filename)
	ActOn(err)

	result = readIntSlice2D(string(content), sep1, sep2)

	if result == nil || len(result) == 0 {
		return nil
	}

	return
}

func readIntSlice2D(content, sep1, sep2 string) (result [][]int) {
	for _, line := range strings.Split(content, sep1) {
		line = strings.TrimSpace(line)
		chunks := strings.Split(line, sep2)
		series := make([]int, len(chunks))
		for i, chunk := range chunks {
			number, err := strconv.Atoi(chunk)
			ActOn(err)

			series[i] = number
		}
		result = append(result, series)
	}

	return
}
