// Package prime implements a solution for the exercise titled `Nth Prime'.
//
// I'd like to write like this someday...
// https://github.com/clojure/clojure-contrib/blob/master/modules/lazy-seqs/src/main/clojure/clojure/contrib/lazy_seqs.clj
//
package prime

var primes []int = []int{2, 3, 5, 7, 11, 13}

// Nth returns nth prime number.
func Nth(n int) (int, bool) {
	divisible := func(n int) bool {
		for _, p := range primes[1:] {
			if n%p == 0 {
				return true
			}
		}
		return false
	}

	if n <= 0 {
		return 0, false
	}
	for len(primes) < n {
		i := primes[len(primes)-1] + 2
		for ; divisible(i); i += 2 {
		}
		primes = append(primes, i)
	}
	return primes[n-1], true
}
