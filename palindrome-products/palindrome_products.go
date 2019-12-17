// Package palindrome implements a solution of the exercise titled `Palindrome Products'.
package palindrome

import (
	"errors"
	"strconv"
)

// Product is the structure returned by the function Products.
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func isPalindromic(n int) bool {
	s := strconv.Itoa(n)
	for i := range s[:len(s)/2] {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// Products returns the smallest and largest palindromes which are products of numbers within that range.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
	} else {
		pmin, pmax = Product{}, Product{}
		for i := fmin; i <= fmax; i++ {
			for j := i; j <= fmax; j++ {
				prod := i * j
				if isPalindromic(prod) {
					if pmin.Product == 0 || prod < pmin.Product {
						pmin = Product{prod, [][2]int{{i, j}}}
					} else if prod == pmin.Product {
						pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
					}
					if pmax.Product == 0 || prod > pmax.Product {
						pmax = Product{prod, [][2]int{{i, j}}}
					} else if prod == pmax.Product {
						pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
					}
				}
			}
		}
		if pmin.Product == 0 {
			err = errors.New("no palindromes")
		}
		if pmin.Product == pmax.Product {
			pmin = Product{}
		}
	}
	return
}

/*
BenchmarkPalindromeProducts-4   	      69	  17402190 ns/op	 3073381 B/op	  409704 allocs/op
*/
