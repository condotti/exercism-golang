// Package twobucket implements a solution of the exercise titled `Two Bucket'.
package twobucket

import (
	"errors"
	"fmt"
)

// Solve returns a solution of the problem or an error if not found.
// loop version.
func Solve(bucketOne, bucketTwo, goal int, startBucket string) (goalBucket string, moves, otherBucket int, err error) {
	type frame struct { // stack frame
		contentOne, contentTwo, moves, state int
	}
	stack := []*frame{}
	push := func(c1, c2, m int) {
		stack = append([]*frame{&frame{c1, c2, m + 1, 0}}, stack...)
	}
	pop := func() {
		stack = stack[1:]
	}

	history := map[string]bool{}
	goalBucket, err = "", nil
	if bucketOne <= 0 || bucketTwo <= 0 || goal <= 0 {
		err = errors.New("invalid bucket size or goal amount")
	}

	switch startBucket {
	case "one":
		push(bucketOne, 0, 1)
	case "two":
		push(0, bucketTwo, 1)
	default:
		err = errors.New("invalid start bucket name")
		return
	}

	for len(stack) > 0 {
		top := stack[0]
		switch top.state {
		case 0: // constraints: 0 <= content <= bucket, avoiding loop
			state := fmt.Sprintf("%d,%d", top.contentOne, top.contentTwo)
			if top.contentOne < 0 || top.contentOne > bucketOne ||
				top.contentTwo < 0 || top.contentTwo > bucketTwo ||
				history[state] {
				pop()
			}
			history[state] = true
		case 1: // solution found
			switch {
			case top.contentOne == goal:
				goalBucket, otherBucket, moves = "one", top.contentTwo, top.moves
				return
			case top.contentTwo == goal:
				goalBucket, otherBucket, moves = "two", top.contentOne, top.moves
				return
			}
		case 2: // pour entire content to another
			push(0, top.contentOne+top.contentTwo, top.moves)
		case 3:
			push(top.contentOne+top.contentTwo, 0, top.moves)
		case 4: // pour a part of content to another
			push(top.contentOne-(bucketTwo-top.contentTwo), bucketTwo, top.moves)
		case 5:
			push(bucketOne, top.contentTwo-(bucketOne-top.contentOne), top.moves)
		case 6: // empty a bucket
			push(0, top.contentTwo, top.moves)
		case 7:
			push(top.contentOne, 0, top.moves)
		case 8: // fill a bucket
			push(bucketOne, top.contentTwo, top.moves)
		case 9:
			push(top.contentOne, bucketTwo, top.moves)
		default:
			pop()
		}
		top.state++
	}
	err = errors.New("no solution found")
	return
}

/*
BenchmarkSolve-4   	    6667	    176961 ns/op	   20096 B/op	    1744 allocs/op
*/
