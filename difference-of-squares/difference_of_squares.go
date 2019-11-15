package diffsquares

// SquareOfSum returns (1 + 2 ... + N)^2
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns 1^2 + 2^2 + ... + N^2
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns SquareOfSum(N) - SumOfSquares(N)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
