// Package reverse implements a solution of the exercise titled `Reverse String'.
package reverse

// Reverse returns a revresed string.
func Reverse(a string) string {
	ar := []rune(a)
	br := make([]rune, len(ar))
	for i := range ar {
		br[i] = ar[len(ar)-i-1]
	}
	return string(br)
}
