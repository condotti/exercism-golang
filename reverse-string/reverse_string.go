// Package reverse implements a solution of the exercise titled `Reverse String'.
package reverse

// Reverse returns a revresed string.
func Reverse(a string) string {
	ar := []rune(a)
	br := []rune{}
	for _, c := range ar {
		br = append([]rune{ c }, br...)
	}
	return string(br)
}
