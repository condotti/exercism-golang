// Package summultiples implements a solution for the exercise titled `Sum Of Multiples'.
package summultiples

// SumMultiples returns the sum of all the unique multiples of particular numbers
// up to but not including that number.
func SumMultiples( limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if func() bool {
			for _, j := range divisors {
				if j != 0 && i%j == 0 {
					return true
				}
			}
			return false
		}() {
			sum += i
		}
	}
	return sum
}
