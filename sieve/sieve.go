// Package sieve implements a solution of he exercise titled `Sieve'.
package sieve

// Sieve returns all primes from 2 up to a given number.
func Sieve(limit int) []int {
	// improved method
	divisible := make([]bool, limit-1) // [2, limit-1]
	prime := []int{}
	for i := range divisible {
		if !divisible[i] {
			p := i + 2
			prime = append(prime, p)
			for j := range divisible[i+1:] {
				if !divisible[i+j+1] {
					n := i + j + 3
					if n%p == 0 {
						divisible[i+j+1] = true
					}
				}
			}
		}
	}
	return prime
}

// BenchmarkSieve-4   	    4148	    301508 ns/op	    5336 B/op	      21 allocs/op
