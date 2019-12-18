// Package linkedlist implements a solution of the exercise titled `Linked List'.
package linkedlist

import "errors"

var ErrEmptyList = errors.New("Empty List")

// Node defines a node of linked list.
type Node struct {
	prev, next *Node
	Val        interface{}
}

// List describes a list.
type List struct {
	first, last *Node
}

// Next returns the next node of a node.
func (e *Node) Next() *Node {
	return e.next
}

// Prev returns the previous node of a node.
func (e *Node) Prev() *Node {
	return e.prev
}

// First returns the first node of a list.
func (l *List) First() *Node {
	return l.first
}

// Last returns the last node of a list.
func (l *List) Last() *Node {
	return l.last
}

// PushBack adds specified value to the tail of a list.
func (l *List) PushBack(v interface{}) {
	n := Node{Val: v, prev: l.last}
	if l.last != nil {
		l.last.next = &n
	}
	l.last = &n
	if l.first == nil {
		l.first = &n
	}
}

// NewList returns a list with specified values.
func NewList(args ...interface{}) *List {
	l := List{}
	for _, data := range args {
		l.PushBack(data)
	}
	return &l
}

// PushFront adds specified value to the front of a list.
func (l *List) PushFront(v interface{}) {
	n := Node{Val: v, next: l.first}
	if l.first != nil {
		l.first.prev = &n
	}
	l.first = &n
	if l.last == nil {
		l.last = &n
	}
}

// PopFront returns value of the head node and removes it from the list.
func (l *List) PopFront() (interface{}, error) {
	if l.first == nil {
		return nil, ErrEmptyList
	}
	n := l.first
	if l.last == n {
		l.last = nil
	}
	l.first = n.next
	if l.first != nil {
		l.first.prev = nil
	}
	return n.Val, nil
}

// PopBack returns value of the tail node and removes it from the list.
func (l *List) PopBack() (interface{}, error) {
	if l.last == nil {
		return nil, ErrEmptyList
	}
	n := l.last
	if l.first == n {
		l.first = nil
	}
	l.last = n.prev
	if l.last != nil {
		l.last.next = nil
	}
	return n.Val, nil
}

// Reverse reverses the list.
func (l *List) Reverse() *List {
	r := List{}
	for data, err := l.PopBack(); err == nil; data, err = l.PopBack() {
		r.PushBack(data)
	}
	l.first, l.last = r.first, r.last
	return l
}
