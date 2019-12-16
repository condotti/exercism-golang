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
func Build(record []Record) (*Node, error) {
	if len(record) == 0 {
		return nil, nil
	}
	sort.Slice(record, func(i, j int) bool { return record[i].ID < record[j].ID })
	tree := make([]*Node, len(record))
	for i, r := range record {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}
		tree[i] = &Node{ID: r.ID}
		if r.ID > 0 {
			tree[r.Parent].Children = append(tree[r.Parent].Children, tree[i])
		}
	}
	return tree[0], nil
}

/*
BenchmarkTwoTree-4       	     100	  11688765 ns/op	 3407954 B/op	  131075 allocs/op
BenchmarkTenTree-4       	     760	   1650846 ns/op	  650017 B/op	   15004 allocs/op
BenchmarkShallowTree-4   	     877	   1479511 ns/op	  788312 B/op	   10024 allocs/op
*/
