// Package minesweeper implemnts a solution of exercise titled `Minesweeper'.
package minesweeper

import (
	"errors"
	"regexp"
)

// Count counts the number of mines adjacent to a square.
func (b *Board) Count() error {

	peek := func(r, c int) int {
		if (*b)[r][c] == '*' {
			return 1
		}
		return 0
	}

	for r, row := range *b {
		// Check if the board is valid.
		if len(row) != len((*b)[0]) {
			return errors.New("unaligned")
		}
		if r == 0 || r == len(*b)-1 {
			if !regexp.MustCompile(`^\+-*\+$`).Match(row) {
				return errors.New("incomplete border")
			}
		} else {
			if !regexp.MustCompile(`^\|[0-9* ]*\|$`).Match(row) {
				return errors.New("unknown characters")
			}
			// Count
			for c := range row {
				if (*b)[r][c] == ' ' {
					count := peek(r-1, c-1) + peek(r-1, c) + peek(r-1, c+1) +
						peek(r, c-1) + peek(r, c+1) +
						peek(r+1, c-1) + peek(r+1, c) + peek(r+1, c+1)
					if count > 0 {
						(*b)[r][c] = '0' + byte(count)
					}
				}
			}
		}
	}
	return nil
}

/*
BenchmarkCount-4   	   31333	     37782 ns/op	   24387 B/op	     376 allocs/op
*/
