// Package tree implements a solution of the exercise titled `Tree Building'
package tree

import "errors"

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

func search(id int, p *Node) *Node {
	if p == nil {
		return nil // not found
	}
	if p.ID == id {
		return p // found
	}
	for _, c := range p.Children {
		r := search(id, c)
		if r != nil {
			return r // found
		}
	}
	return nil // not found
}

// Build builds a tree represening the structiure of Record r.
func Build(r []Record) (*Node, error) {
	if len(r) == 0 {
		return nil, nil
	}
	// path 1: sort r by direct insertion
	sorted := make([]int, len(r))
	seen := make([]bool, len(r))
	sorted[0] = -1 // sentinel
	for i := range r {
		if r[i].ID >= len(r) {
			return nil, errors.New("no root node")
		}
		if seen[r[i].ID] {
			return nil, errors.New("duplicate node")
		}
		sorted[r[i].ID] = r[i].Parent
		seen[r[i].ID] = true
	}
	if sorted[0] != 0 {
		return nil, errors.New("root node has parent")
	}
	// path2: build a tree
	root := new(Node)
	for i, j := range sorted {
		if i != 0 {
			target := search(j, root)
			if target == nil {
				return nil, errors.New("non-continuous")
			}
			p := new(Node)
			p.ID = i
			target.Children = append(target.Children, p)
		}
	}
	return root, nil
}
