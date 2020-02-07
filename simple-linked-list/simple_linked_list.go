// Package linkedlist implements a solution of the exercise titled `Simple Linked List'.
package linkedlist

import (
	"errors"
)

// Element represents an element of singly linked list.
type Element struct {
	data int
	next *Element
}

// List represents a list data structure.
type List struct {
	head *Element
	size int
}

// New is the ctor of List
func New(data []int) *List {
	var head *Element
	for i := range data {
		head = &Element{data[len(data)-i-1], head}
	}
	return &List{head, len(data)}
}

// Size returns the length of the list.
func (list *List) Size() int {
	return list.size
}

// Push adds an element into the tail of the list.
func (list *List) Push(datum int) {
	if list.head == nil {
		list.head = &Element{data: datum}
	} else {
		tail := list.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = &Element{data: datum}
	}
	list.size++
}

// Pop removes and returns the tail of the list, or an error when list is empty.
func (list *List) Pop() (int, error) {
	if list.size <= 0 {
		return 0, errors.New("pop from empty list")
	}
	var datum int
	if list.head.next == nil {
		datum, list.head = list.head.data, nil

	} else {
		tail := list.head
		for tail.next != nil {
			tail = tail.next
		}
		prevToTail := list.head
		for prevToTail.next != tail {
			prevToTail = prevToTail.next
		}
		datum, prevToTail.next = tail.data, nil
	}
	list.size--
	return datum, nil
}

// Array converts list into an array.
func (list *List) Array() []int {
	array := make([]int, list.size)
	for e, i := list.head, 0; e != nil; e, i = e.next, i+1 {
		array[i] = e.data
	}
	return array
}

// Reverse returns a list of reversed order of the list.
func (list *List) Reverse() *List {
	data := list.Array()
	reversed := make([]int, list.size)
	for i := 0; i < len(data); i++ {
		reversed[i] = data[len(data)-i-1]
	}
	return New(reversed)
}

/*
BenchmarkNewList-4       	 2837767	       388 ns/op	     176 B/op	      11 allocs/op
BenchmarkListSize-4      	1000000000	         0.407 ns/op	       0 B/op	       0 allocs/op
BenchmarkListPush-4      	    1431	    782311 ns/op	   16000 B/op	    1000 allocs/op
BenchmarkListPop-4       	     890	   1319807 ns/op	       0 B/op	       0 allocs/op
BenchmarkListToArray-4   	27323772	        52.2 ns/op	      80 B/op	       1 allocs/op
BenchmarkListReverse-4   	 2465592	       453 ns/op	     336 B/op	      13 allocs/op
*/
