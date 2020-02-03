// Package queenattack implements a solution of the exercise titled `Queen Attack'.
package queenattack

import "errors"

// CanQueenAttack receives 2 positions on the board, returns if the queen can
// attack, or signals if positions are incorrect.
func CanQueenAttack(black, white string) (attack bool, err error) {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	if len(black) == 2 && len(white) == 2 {
		ab, aw := []rune(black), []rune(white)
		b, w := []int{int(ab[0] - 'a'), int(ab[1] - '1')}, []int{int(aw[0] - 'a'), int(aw[1] - '1')}
		if 0 <= b[0] && b[0] < 8 && 0 <= b[1] && b[1] < 8 &&
			0 <= w[0] && w[0] < 8 && 0 <= w[1] && w[1] < 8 &&
			!(b[0] == w[0] && b[1] == w[1]) {
			return b[0] == w[0] || b[1] == w[1] ||
				abs(b[0]-w[0]) == abs(b[1]-w[1]), nil
		}
	}
	return false, errors.New("invalid input")
}

/*
BenchmarkCanQueenAttack-4   	 1998660	       602 ns/op	      96 B/op	       6 allocs/op
*/
