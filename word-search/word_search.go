// Package wordsearch implements a solution of the exercise titled `Word Search'.
package wordsearch

import (
	"errors"
)

// Solve returns the first and last locations of given list of words.
func Solve(words, puzzle []string) (map[string][2][2]int, error) {

	dir := [][2]int{
		{1, 0},   // left to right
		{0, -1},  // top to bottom
		{-1, 0},  // right to left
		{0, 1},   // bottom to top
		{1, -1},  // upper-left to lower-right
		{-1, -1}, // upper-right to lower-left
		{1, 1},   // lower-left to upper-right
		{-1, 1},  // lower-right to upper-left
	}

	mtx := [][]rune{}
	for _, line := range puzzle {
		mtx = append(mtx, []rune(line))
	}

	match := func(word []rune, x, y int) (pos [2][2]int, found bool) {
		found = false
		if mtx[y][x] != word[0] {
			return
		}
		pos[0][0], pos[0][1] = x, y
		for _, d := range dir {
			nx, ny := x+d[0], y+d[1]
			found = true
			for _, x := range word[1:] {
				if ny >= 0 && ny < len(mtx) &&
					nx >= 0 && nx < len(mtx[0]) &&
					mtx[ny][nx] == x {
					pos[1][0], pos[1][1] = nx, ny
					nx, ny = nx+d[0], ny+d[1]
					continue
				}
				found = false
				break
			}
			if found {
				return
			}
		}
		return
	}

	solution := map[string][2][2]int{}
	for _, word := range words {
		for y, row := range mtx {
			for x := range row {
				if pos, found := match([]rune(word), x, y); found {
					solution[word] = pos
				}
			}
		}
	}
	if len(solution) != len(words) {
		return nil, errors.New("word not in the puzzle")
	}
	return solution, nil
}

/*
BenchmarkSolve-4   	    3852	    280962 ns/op	  164176 B/op	    5253 allocs/op
*/
