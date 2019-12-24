// Package twobucket implements a solution of the exercise titled `Two Bucket'.
package twobucket

import (
	"errors"
	"fmt"
)

// Solve returns a solution of the problem or an error if not found.
// recursive version.
func Solve(bucketOne, bucketTwo, goal int, startBucket string) (goalBucket string, moves, otherBucket int, err error) {
	history := map[string]bool{}
	var step func(int, int, int)
	step = func(contentOne, contentTwo, n int) {
		// constraints: 0 <= content <= bucket, avoiding loop
		state := fmt.Sprintf("%d,%d", contentOne, contentTwo)
		if contentOne < 0 || contentOne > bucketOne || contentTwo < 0 || contentTwo > bucketTwo || history[state] {
			return
		}
		// solution found
		if contentOne == goal {
			goalBucket, otherBucket, moves = "one", contentTwo, n
			return
		}
		if contentTwo == goal {
			goalBucket, otherBucket, moves = "two", contentOne, n
			return
		}
		history[state] = true
		// 1. pour one bucket to another
		//   1.1 entire content to another
		step(0, contentOne+contentTwo, n+1)
		step(contentOne+contentTwo, 0, n+1)
		//   1.2 a part of content to another
		step(contentOne-bucketTwo+contentTwo, bucketTwo, n+1)
		step(bucketOne, contentTwo-bucketOne+contentOne, n+1)
		// 2. empty a bucket
		step(0, contentTwo, n+1)
		step(contentOne, 0, n+1)
		// 3. fill a bucket
		step(bucketOne, contentTwo, n+1)
		step(contentOne, bucketTwo, n+1)
		return
	}
	goalBucket, err = "", nil
	if bucketOne <= 0 || bucketTwo <= 0 || goal <= 0 {
		err = errors.New("invalid bucket size or goal amount")
	}
	switch startBucket {
	case "one":
		step(bucketOne, 0, 1)
	case "two":
		step(0, bucketTwo, 1)
	default:
		err = errors.New("invalid start bucket name")
	}
	if goalBucket == "" {
		err = errors.New("no solution found")
	}
	return
}

/*
BenchmarkSolve-4   	    6667	    176961 ns/op	   20096 B/op	    1744 allocs/op
*/
