// Package tree implements a solution of the exercise titled `Tree Building'
package tree

import (
	"errors"
	"sort"
)

// Record defines input tree format
type Record struct {
	ID     int
	Parent int
}

// Node defines output tree format
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree represening the structiure of Record r.
func Build(r []Record) (*Node, error) {
	if len(r) == 0 {
		return nil, nil
	}
	sort.Slice(r, func(i, j int) bool { return r[i].ID < r[j].ID })
	if r[0].ID != 0 || r[0].Parent != 0 {
		return nil, errors.New("invalid input")
	}
	tree := make([]*Node, len(r))
	tree[0] = &Node{}
	for i := 1; i < len(r); i++ {
		if i == r[i].ID && i > r[i].Parent {
			tree[i] = &Node{ID: i}
			tree[r[i].Parent].Children = append(tree[r[i].Parent].Children, tree[i])
		} else {
			return nil, errors.New("invalid input")
		}
	}
	return tree[0], nil
}

/*
BenchmarkTwoTree-4       	      97	  11834644 ns/op	 3407956 B/op	  131075 allocs/op
BenchmarkTenTree-4       	     765	   1659616 ns/op	  650018 B/op	   15004 allocs/op
BenchmarkShallowTree-4   	     859	   1357301 ns/op	  788312 B/op	   10024 allocs/op
*/
