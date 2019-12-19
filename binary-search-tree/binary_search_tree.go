// Package binarysearchtree implements a solution of the exercise titled `Binary Search Tree'.
package binarysearchtree

// SearchTreeData defines a tree node of the binary search tree.
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst returns a tree node which has specified data.
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

// Insert is a method to insert a tree node to the tree.
func (t *SearchTreeData) Insert(data int) {
	if data > t.data {
		if t.right == nil {
			t.right = Bst(data)
		} else {
			t.right.Insert(data)
		}
	} else {
		if t.left == nil {
			t.left = Bst(data)
		} else {
			t.left.Insert(data)
		}
	}
}

// MapString returns a slice of strings which is formatted by the function.
func (t *SearchTreeData) MapString(f func(int) string) []string {
	out := []string{}
	if t.left != nil {
		out = t.left.MapString(f)
	}
	out = append(out, f(t.data))
	if t.right != nil {
		out = append(out, t.right.MapString(f)...)
	}
	return out
}

// MapInt returns s slice of the tree data which is applied by the function.
func (t *SearchTreeData) MapInt(f func(int) int) []int {
	out := []int{}
	if t.left != nil {
		out = t.left.MapInt(f)
	}
	out = append(out, f(t.data))
	if t.right != nil {
		out = append(out, t.right.MapInt(f)...)
	}
	return out
}
