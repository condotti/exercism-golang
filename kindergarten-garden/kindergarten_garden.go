// Package kindergarten implements a solution of the exercise titiled `Kindergarten Garden'.
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden represents a collection of a child and his/her own grasses.
type Garden map[string][]string

var plants = map[byte]string{'G': "grass", 'C': "clover", 'R': "radishes", 'V': "violets"}

// NewGarden is a ctor of Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	for _, c := range diagram {
		switch c {
		case 'G', 'C', 'R', 'V', '\n':
			continue
		default:
			return nil, errors.New("invalid cup code")
		}
	}
	sorted := make([]string, cap(children))
	copy(sorted, children)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	g := Garden{}
	rows := strings.Split(diagram, "\n")[1:]
	if len(rows) != 2 || len(rows[0]) != len(rows[1]) || len(rows[0])&1 == 1 {
		return nil, errors.New("invalid input")
	}
	for _, child := range sorted {
		if _, found := g[child]; found {
			return nil, errors.New("duplicate name")
		}
		g[child] = []string{plants[rows[0][0]], plants[rows[0][1]], plants[rows[1][0]], plants[rows[1][1]]}
		rows[0], rows[1] = rows[0][2:], rows[1][2:]
	}
	return &g, nil
}

// Plants returns the plants which the child owns.
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := map[string][]string(*g)[child]
	return plants, ok
}

/*
BenchmarkNewGarden-4       	   96256	     12029 ns/op	    7356 B/op	      99 allocs/op
BenchmarkGarden_Plants-4   	 6137862	       214 ns/op	       0 B/op	       0 allocs/op
*/
