package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadNumbers(filename string) []int {
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

	if len(result) == 0 {
		return nil
	}

	return result
}

type Instruction struct {
	Direction string
	Distance  int
}

// ReadInstructions will parse file provided in format:
//
// stringKey1 intValue1\n
// stringKey2 intValue2\n
func ReadInstructions(filename string) []Instruction {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	result := make([]Instruction, 0)

	for _, line := range strings.Split(string(content), "\n") {
		chunks := strings.Split(line, " ")
		distance, err := strconv.Atoi(chunks[1])

		if err != nil {
			continue
		}

		result = append(result, Instruction{chunks[0], distance})
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
