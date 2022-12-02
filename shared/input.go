package shared

import (
	"os"
	"reflect"
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

func ReadSlice2D[T int | rune](filename, sep1, sep2 string) (result [][]T) {
	content, err := os.ReadFile(filename)
	ActOn(err)

	result = readSlice2D[T](string(content), sep1, sep2)

	if result == nil || len(result) == 0 {
		return nil
	}

	return
}

func readSlice2D[T int | rune](content, sep1, sep2 string) (result [][]T) {
	for _, line := range strings.Split(content, sep1) {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			chunks := strings.Split(line, sep2)
			series := make([]T, len(chunks))

			switch reflect.TypeOf((*T)(nil)).Elem() {
			case reflect.TypeOf((*int)(nil)).Elem():
				for i, chunk := range chunks {
					number, err := strconv.Atoi(chunk)
					ActOn(err)

					series[i] = T(number)
				}
			case reflect.TypeOf((*rune)(nil)).Elem():
				for i, chunk := range chunks {
					symbol := []rune(chunk)

					series[i] = T(symbol[0])
				}
			}

			result = append(result, series)
		}
	}

	return
}

func identity[T any](v T) T {
	return v
}
