package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntFromLines(filename string) []int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	result := make([]int, 0)

	for _, line := range strings.Split(string(content), "\n") {
		number, err := strconv.Atoi(line)

		if err != nil {
			continue
		}

		result = append(result, number)
	}

	return result
}
