// Package matrix implements a solution of the exercise titled `Saddle Points'.
package matrix

import (
	"math"
)

// Pair represents a pair of int.
type Pair [2]int

func max(xs []int) int {
	maximum := math.MinInt32
	for _, x := range xs {
		if x > maximum {
			maximum = x
		}
	}
	return maximum
}

func min(xs []int) int {
	minimum := math.MaxInt32
	for _, x := range xs {
		if x < minimum {
			minimum = x
		}
	}
	return minimum
}

// Saddle returns saddle points of the matrix.
func (m *Matrix) Saddle() []Pair {
	sp := []Pair{}
	rows := m.Rows()
	cols := m.Cols()
	for r, row := range rows {
		for c, col := range cols {
			if (*m)[r][c] == max(row) && (*m)[r][c] == min(col) {
				sp = append(sp, Pair{r, c})
			}
		}
	}
	return sp
}

/*
BenchmarkSaddle-4   	  300548	      3992 ns/op	    2208 B/op	      68 allocs/op
*/
