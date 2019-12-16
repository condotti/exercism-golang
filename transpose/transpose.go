// Package transpose implements a solution of the exercise titled `Transpose'.
package transpose

// Transpose returns a transposed text.
func Transpose(input []string) []string {
	output := []string{}
	for r, line := range input {
		for c, letter := range line {
			if c >= len(output) {
				output = append(output, "")
			}
			for ; len(output[c]) < r; output[c] += " " {
			}
			output[c] += string(letter)
		}
	}
	return output
}

/*
BenchmarkTranspose-4   	   94830	     12526 ns/op	    3856 B/op	     183 allocs/op
*/
