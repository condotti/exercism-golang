// Package grains implements a solution of the exercise titled `Grains'.
package grains

import "errors"

// Square returns the number of grains on Nth square
func Square(n int) (uint64, error) {
	// n Square(n) = 2^(n-1)
	// 1        1  = 2^0 = 0b0...0001
	// 2        2  = 2^1 = 0b0...0010
	// 3        4  = 2^2 = 0b0...0100
	// 4        8  = 2^3 = 0b0...1000
	// :        :     :        :
	// 64           2^63 = 0b1...0000
	if 0 < n && n <= 64 {
		return uint64(1) << (n - 1), nil
	}
	return 0, errors.New("input must be positive")
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	// As a comment above, Square(1)+Square(2)+...+Square(64) = 0b1...111 (64bits).
	return ^uint64(0)
}
