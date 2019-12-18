// Package binarysearch implements a solution of the exercise titled `Binary Search`.
package binarysearch

// SearchInts returns index of the key in the slice.
func SearchInts(slice []int, key int) int {
	if len(slice) != 0 {
		for s, e, x := 0, len(slice), len(slice)/2; ; {
			px := x
			switch {
			case slice[x] == key:
				return x
			case key > slice[x]:
				x, s = (x+e)/2, x
			case key < slice[x]:
				x, e = (s+x)/2, x
			}
			if px == x {
				break
			}
		}
	}
	return -1
}

/*
Benchmark1e2-4   	144100714	         8.69 ns/op	       0 B/op	       0 allocs/op
Benchmark1e4-4   	70841179	        23.4 ns/op	       0 B/op	       0 allocs/op
Benchmark1e6-4   	34339872	        33.9 ns/op	       0 B/op	       0 allocs/op
Benchmark1e8-4   	30827882	        44.4 ns/op	       0 B/op	       0 allocs/op
*/
