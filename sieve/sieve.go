// Package sieve implements a solution of he exercise titled `Sieve'.
package sieve

// Sieve returns all primes from 2 up to a given number.
func Sieve(limit int) []int {
	// naive method: seems not efficient
	number := []int{}
	for i := 2; i <= limit; i++ {
		number = append(number, i)
	}
	prime := []int{}
	for len(number) > 0 {
		p := number[0]
		number = number[1:]
		tmp := []int{}
		for _, n := range number {
			if n%p != 0 {
				tmp = append(tmp, n)
			}
		}
		number = tmp
		prime = append(prime, p)
	}
	return prime
}

// BenchmarkSieve-4   	    4010	    291001 ns/op	  366968 B/op	    1309 allocs/op
