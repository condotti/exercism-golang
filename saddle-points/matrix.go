// Package matrix implements a solution of the exercise titled `Matrix'
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents 2dimensional matrix
type Matrix [][]int

// New is ctor of Matrix
func New(s string) (*Matrix, error) {
	mtx := Matrix{}
	for i, r := range strings.Split(s, "\n") {
		row := []int{}
		for _, c := range strings.Split(strings.Trim(r, " "), " ") {
			v, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			row = append(row, v)
		}
		if i != 0 && len(row) != len(mtx[0]) {
			return nil, errors.New("uneven rows")
		}
		mtx = append(mtx, row)
	}
	return &mtx, nil
}

// Rows returns the list of Matrix's rows
func (m *Matrix) Rows() [][]int {
	rows := [][]int{}
	for _, r := range *m {
		row := make([]int, len(r))
		copy(row, r)
		rows = append(rows, row)
	}
	return rows
}

// Cols returns the list of Matrix's columns
func (m *Matrix) Cols() [][]int {
	cols := [][]int{}
	for i := range (*m)[0] {
		col := []int{}
		for j := 0; j < len(*m); j++ {
			col = append(col, (*m)[j][i])
		}
		cols = append(cols, col)
	}
	return cols
}

// Set sets the value of specified row and column of Matrix
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || r >= len(*m) || c < 0 || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = val
	return true
}
