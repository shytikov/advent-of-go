package utils

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

// ReadIntSliceFromRuneSlice will read file and will treat each line as slice of runes
// Each rune then will be converted to integer
func ReadIntSliceFromRuneSlice(filename string) (result [][]int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	for _, line := range strings.Split(string(content), "\n") {
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

	if len(result) == 0 {
		return nil
	}

	return
}

func ReadIntSliceFromLine(filename string) (result [][]int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	for _, line := range strings.Split(string(content), "\n") {
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

	if len(result) == 0 {
		return nil
	}

	return
}
