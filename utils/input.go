package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("cannot find input file"))
	}

	return strings.Split(string(content), "\n")
}
