package local

import (
	"github.com/shytikov/advent-of-go/shared"
	"os"
	"strconv"
	"strings"
)

type Data []Command

// Command is type representing a pair of the name and argument
type Command struct {
	Name     string
	Argument int
}

// Read will parse file provided in format:
//
// stringKey1 intValue1\n
// stringKey2 intValue2\n
func Read(filename string) (result Data) {
	content, err := os.ReadFile(filename)
	shared.ActOn(err)

	for _, line := range strings.Split(string(content), "\n") {
		chunks := strings.Split(line, " ")
		distance, err := strconv.Atoi(chunks[1])

		if err != nil {
			continue
		}

		result = append(result, Command{chunks[0], distance})
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
