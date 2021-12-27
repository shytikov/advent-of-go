package shared

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntFromLine(filename string) (result []int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	for _, line := range strings.Split(string(content), "\n") {
		number, err := strconv.Atoi(line)

		if err != nil {
			continue
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
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

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
		serie := make([]int, len(chars))
		for i, char := range chars {
			number, err := strconv.Atoi(string([]rune{char}))

			if err != nil {
				return nil
			}

			serie[i] = number
		}

		result = append(result, serie)
	}

	return
}

func ReadRuneSlicesFromLines(filename string) (result [][]rune) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

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

func ReadIntSlicesFromLines(filename string) (result [][]int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	result = readIntSlicesFromLines(string(content))

	if result == nil || len(result) == 0 {
		return nil
	}

	return
}

func readIntSlicesFromLines(content string) (result [][]int) {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		chunks := strings.Split(line, ",")
		serie := make([]int, len(chunks))
		for i, chunk := range chunks {
			number, err := strconv.Atoi(chunk)

			if err != nil {
				return nil
			}

			serie[i] = number
		}
		result = append(result, serie)
	}

	return
}
