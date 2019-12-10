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
		pmin, pmax, empty = Product{}, Product{}, Product{}
		for i := fmin; i <= fmax; i++ {
			for j := i; j <= fmax; j++ {
				prod := i * j
				if isPalindromic(prod) {
					if pmin == empty || prod < pmin.Product {
						pmin.Product = Product{prod, [][2]int{i, j}}
					} else if prod == pmin.Product {
						pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
					}
					if pmax == empty || prod > pmax.Product {
						pmax.Product = Product{prod, [][2]int{i, j}}
					} else if prod == pman.Product {
						pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
					}
				}
			}
		}
		if pmin == empty {
			err = errors.New("no palindromes")
		}
		if pmin.Product == pmax.Product {
			pmin = empty
		}
		return
	}
}

/*
BenchmarkTwoTree-4       	     100	  11633995 ns/op	 3407976 B/op	  131075 allocs/op
BenchmarkTenTree-4       	     780	   1662342 ns/op	  650017 B/op	   15004 allocs/op
BenchmarkShallowTree-4   	     871	   1359794 ns/op	  788312 B/op	   10024 allocs/op
*/
