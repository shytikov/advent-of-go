package local

import (
	"strings"

	"github.com/shytikov/advent-of-go/shared"
)

type Track struct {
	path    string
	visited map[string]int
}

func (t Track) String() string {
	if strings.HasSuffix(t.path, ",") {
		return t.path[0 : len(t.path)-1]
	} else {
		return t.path
	}
}

func (t Track) GetStats() (start, end, max_big, max_small int) {
	for name, count := range t.visited {
		if name == "start" {
			start += count
		} else if name == "end" {
			end += count
		} else {
			if strings.ToUpper(name) == name {
				if max_big < count {
					max_big = count
				}
			} else {
				if max_small < count {
					max_small = count
				}
			}
		}
	}

	return
}

func (t *Track) Add(current *shared.Node) string {
	name := current.Value.(string)

	t.path = t.path + name + ","

	t.visited = make(map[string]int)
	for _, name := range strings.Split(t.String(), ",") {
		if value, ok := t.visited[name]; ok {
			t.visited[name] = value + 1
		} else {
			t.visited[name] = 1
		}
	}

	return t.String()
}

func (t Track) IsComplete() bool {
	return strings.Contains(t.path, "end")
}

func (t Track) IsCircular() bool {
	return strings.Contains(t.path, ",start")
}
