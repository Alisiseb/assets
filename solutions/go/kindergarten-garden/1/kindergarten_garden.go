package kindergarten

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

// Define the Garden type here.
type Garden map[string][]string

var plantsNames = map[string]string{
	"G": "grass",
	"C": "clover",
	"R": "radishes",
	"V": "violets",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	diagramSlice := strings.Fields(diagram)
	if !diagramIsValid(diagram) {
		return nil, errors.New("invalid diagram format")
	}
	class := make(Garden)
	childs:= make([]string, len(children))
	copy(childs, children)
	slices.Sort(childs)
	if !IsChildrenValid(childs) {
		return nil, errors.New("duplicate names in children list")
	}
	for i, v := range childs {
		plants := []string{}
		for j := 0; j < 2; j++ {
			for p := 2 * i; p < 2*i+2; p++ {
				plantName, ok := plantsNames[string(diagramSlice[j][p])]
				if !ok {
					return nil, errors.New("invalid plant code in diagram")
				}
				plants = append(plants, plantName)
			}
		}

		class[v] = plants

	}
	return &class, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	result, ok := (*g)[child]
	return result, ok
}
func diagramIsValid(diagram string) bool {
	diagramSlice := strings.Fields(diagram)
	if len(diagramSlice) != 2 || len(diagramSlice[0])%2 != 0 || len(diagramSlice[0]) != len(diagramSlice[1]) {
		return false
	}
	re := regexp.MustCompile(`^\n([GCRV]{2})+\n([GCRV]{2})+$`)
	return re.MatchString(diagram)
}
func IsChildrenValid(v []string) bool {
	for i := 0; i < len(v); i++ {
		if i+1 < len(v) && v[i] == v[i+1] {
			return false
		}
	}
	return true
}
