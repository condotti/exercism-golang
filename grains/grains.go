package grains

import "errors"

// Square returns the number of grains on Nth square
func Square(n int) (uint64, error) {
	if 0 < n && n <= 64 {
		return uint64(1) << (n - 1), nil
	}
	return 0, errors.New("input must be positive")
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return ^uint64(0)
}
