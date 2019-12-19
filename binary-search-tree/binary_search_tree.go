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
	target := t
	for t != nil {
		target = t
		if data > t.data {
			t = t.right
		} else {
			t = t.left
		}
	}
	if data > target.data {
		target.right = Bst(data)
	} else {
		target.left = Bst(data)
	}
}

// MapString returns a slice of strings which is formatted by the function.
func (t *SearchTreeData) MapString(f func(int) string) []string {
	out := []string{}
	stack := []*SearchTreeData{t}
	state := map[*SearchTreeData]int{}
	for len(stack) > 0 {
		switch state[stack[0]] {
		case 0:
			state[stack[0]]++
			if stack[0].left != nil {
				stack = append([]*SearchTreeData{stack[0].left}, stack...)
			}
		case 1:
			state[stack[0]]++
			out = append(out, f(stack[0].data))
		case 2:
			state[stack[0]]++
			if stack[0].right != nil {
				stack = append([]*SearchTreeData{stack[0].right}, stack...)
			}
		default:
			stack = stack[1:]
		}
	}
	return out
}

// MapInt returns s slice of the tree data which is applied by the function.
func (t *SearchTreeData) MapInt(f func(int) int) []int {
	out := []int{}
	stack := []*SearchTreeData{t}
	state := map[*SearchTreeData]int{}
	for len(stack) > 0 {
		switch state[stack[0]] {
		case 0:
			state[stack[0]]++
			if stack[0].left != nil {
				stack = append([]*SearchTreeData{stack[0].left}, stack...)
			}
		case 1:
			state[stack[0]]++
			out = append(out, f(stack[0].data))
		case 2:
			state[stack[0]]++
			if stack[0].right != nil {
				stack = append([]*SearchTreeData{stack[0].right}, stack...)
			}
		default:
			stack = stack[1:]
		}
	}
	return out
}
