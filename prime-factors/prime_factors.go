// Package prime implements a solution of the exercise titled `Prime Factors'.
package prime

// Factors returns a list of prime factors of a given natural number.
func Factors(n int64) []int64 {
	c := int64(2)
	factors := []int64{}
	for n > 1 {
		if n%c == 0 {
			n /= c
			factors = append(factors, c)
		} else {
			c++
		}
	}
	return factors
}
