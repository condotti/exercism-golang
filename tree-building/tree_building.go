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
	tree := make([]*Node, len(r))
	for i, inserting := range r {
		tree[i] = &Node{ID: inserting.ID}
		if inserting.ID != i {
			return nil, errors.New("non-continuous")
		}
		if inserting.ID == 0 {
			if inserting.Parent != 0 {
				return nil, errors.New("root node has parent")
			}
		} else {
			if inserting.ID <= inserting.Parent {
				return nil, errors.New("cycle detected")
			}
			tree[inserting.Parent].Children = append(tree[inserting.Parent].Children, tree[i])
		}
	}
	return tree[0], nil
}

/*
BenchmarkTwoTree-4       	      99	  11595263 ns/op	 3407959 B/op	  131075 allocs/op
BenchmarkTenTree-4       	     760	   1628585 ns/op	  650016 B/op	   15004 allocs/op
BenchmarkShallowTree-4   	     829	   1370239 ns/op	  788312 B/op	   10024 allocs/op
*/
