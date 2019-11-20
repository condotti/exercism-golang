// Package lsproduct implements a solution of the exercise titled `Largest Series Product'.
package lsproduct

import "errors"

// LargestSeriesProduct returns the  largest product for a contiguous substring of
// digits of length n, or an error.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span >= 0 && span <= len(digits) {
		product := 0
		d := []rune(digits)
		for i := 0; i < len(d)-span+1; i++ {
			p := 1
			for j := 0; j < span; j++ {
				c := int(d[i+j] - '0')
				if 0 <= c && c <= 9 {
					p *= c
				} else {
					return -1, errors.New("non-digit input")
				}
			}
			if p > product {
				product = p
			}
		}
		return product, nil
	}
	return -1, errors.New("invalid span or span too large")
}
