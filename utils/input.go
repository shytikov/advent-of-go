package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntsFromLines(filename string) (result []int) {
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

	return result
}

func ReadIntSlicesFromLines(filename string) (result [][]int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	for _, line := range strings.Split(string(content), "\n") {
		serie := make([]int, 0)
		for _, char := range []rune(line) {
			number, err := strconv.Atoi(string([]rune{char}))

			if err != nil {
				return nil
			}

			serie = append(serie, number)
		}

		result = append(result, serie)
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
