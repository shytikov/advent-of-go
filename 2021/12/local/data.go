package local

import (
	"io/ioutil"
	"strings"

	"github.com/shytikov/advent-of-go/shared"
)

type Data *shared.Node

func Read(filename string) (result Data) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	return read(string(content))
}

func read(content string) (result Data) {
	lines := strings.Split(content, "\n")

	caves := make(map[string]*shared.Node)
	caves["start"] = &shared.Node{
		Value: "start",
		Links: make([]*shared.Node, 0),
	}

	caves["end"] = &shared.Node{
		Value: "end",
		Links: make([]*shared.Node, 0),
	}

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) > 0 {
			names := strings.Split(line, "-")
			temp := make([]*shared.Node, 2)

			for i, name := range names {
				if cave, ok := caves[name]; ok {
					temp[i] = cave
				} else {
					caves[name] = &shared.Node{
						Value: name,
						Links: make([]*shared.Node, 0),
					}
					temp[i] = caves[name]
				}
			}

			temp[0].Links = append(temp[0].Links, temp[1])
			temp[1].Links = append(temp[1].Links, temp[0])
		}
	}

	return caves["start"]
}
