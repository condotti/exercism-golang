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
	nM := [][]int{}
	for i, r := range strings.Split(s, "\n") {
		nR := []int{}
		for _, c := range strings.Split(strings.Trim(r, " "), " ") {
			x, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			nR = append(nR, x)
		}
		if i != 0 && len(nR) != len(nM[0]) {
			return nil, errors.New("uneven rows")
		}
		nM = append(nM, nR)
	}
	m := Matrix(nM)
	return &m, nil
}

// Rows returns the rows list
func (m *Matrix) Rows() [][]int {
	n := [][]int(*m)
	nM := [][]int{}
	for i := 0; i < len(n); i++ {
		nR := []int{}
		for j := 0; j < len(n[i]); j++ {
			nR = append(nR, n[i][j])
		}
		nM = append(nM, nR)
	}
	return nM
}

// Cols returns the columns list
func (m *Matrix) Cols() [][]int {
	n := [][]int(*m)
	nM := [][]int{}
	for i := 0; i < len(n[0]); i++ {
		nR := []int{}
		for j := 0; j < len(n); j++ {
			nR = append(nR, n[j][i])
		}
		nM = append(nM, nR)
	}
	return nM
}

// Set sets the value of specified row and column of Matrix
func (m *Matrix) Set(r, c, val int) bool {
	n := [][]int(*m)
	if r < 0 || r >= len(n) || c < 0 || c >= len(n[0]) {
		return false
	}
	n[r][c] = val
	return true
}
