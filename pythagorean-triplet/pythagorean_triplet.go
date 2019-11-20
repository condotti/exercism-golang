// Package pythagorean implements a solution of the exercise titiled `Pythagorean Triplet'.
package pythagorean

// Triplet represents a pythagorean triplet.
type Triplet [3]int

func isPythagorean(a, b, c int) bool {
	return a*a+b*b == c*c || b*b+c*c == a*a || c*c+a*a == b*b
}

// Range returns pythagorean triplets ranging in min <= a, b, c <= max.
func Range(min, max int) []Triplet {
	pythagorean := []Triplet{}
	// to ensure a < b < c; fix a and c, then increment b from a to c (exclusive)
	for a := min; a <= max-2; a++ {
		for c := a + 2; c <= max; c++ {
			for b := a + 1; b < c; b++ {
				if isPythagorean(a, b, c) {
					pythagorean = append(pythagorean, Triplet{a, b, c})
				}
			}
		}
	}
	return pythagorean
}

// Sum returns pythagorean triplets where a + b + c = n.
func Sum(n int) []Triplet {
	pythagorean := []Triplet{}
	for a := 1; a < n/3; a++ {
		// the condition 2*b < n-a is from b < c === b < n-a-b
		for b := a + 1; 2*b < n-a; b++ {
			if isPythagorean(a, b, n-a-b) {
				pythagorean = append(pythagorean, Triplet{a, b, n - a - b})
			}
		}
	}
	return pythagorean
}
