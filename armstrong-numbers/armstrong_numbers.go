// Package armstrong implements a solution of the exercise titled `Armstrong Numbers'
package armstrong

// IsNumber determines if a given number is an armstrong number or not.
func IsNumber(n int) bool {
	digits := func(n int) []int {
		ds := []int{}
		for n > 0 {
			ds = append(ds, n%10)
			n /= 10
		}
		return ds

	}(n)
	sum := 0
	for _, digit := range digits {
		pow := 1
		for range digits {
			pow *= digit
		}
		sum += pow
	}
	return sum == n
}
