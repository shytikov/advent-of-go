package local

import (
	"fmt"
	"github.com/shytikov/advent-of-go/shared"
	"io/ioutil"
	"strings"
)

type Data struct {
	Points []shared.Point
	Folds  []int
}

func Read(filename string) (result Data) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	return parseData(string(content))
}

func parseData(content string) (result Data) {
	chunks := strings.Split(trimLines(content), "\n\n")

	fmt.Println(len(chunks))

	return
}

func trimLines(content string) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return strings.Join(lines, "\n")
}
