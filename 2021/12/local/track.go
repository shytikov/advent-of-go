package local

import (
	"strings"

	"github.com/shytikov/advent-of-go/shared"
)

type Track struct {
	path    string
	Visited map[string]int
}

func (t Track) String() string {
	if strings.HasSuffix(t.path, ",") {
		return t.path[0 : len(t.path)-1]
	} else {
		return t.path
	}
}

func (t *Track) Add(current *shared.Node) string {
	name := current.Value.(string)

	t.path = t.path + name + ","

	t.Visited = make(map[string]int)
	for _, name := range strings.Split(t.String(), ",") {
		if value, ok := t.Visited[name]; ok {
			t.Visited[name] = value + 1
		} else {
			t.Visited[name] = 1
		}
	}

	return t.String()
}
