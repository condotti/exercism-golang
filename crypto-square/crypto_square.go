// Package cryptosquare implements a solution of the exercise titled `Crypto Square'.
package cryptosquare

import "math"

// Encode converts a plain text specieid into the cipher text.
func Encode(plainText string) string {
	normalized := []rune{}
	for _, c := range plainText {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			normalized = append(normalized, c)
		} else if 'A' <= c && c <= 'Z' {
			normalized = append(normalized, c-'A'+'a')
		}
	}
	c := int(math.Ceil(math.Sqrt(float64(len(normalized)))))
	r := int(math.Ceil(float64(len(normalized)) / float64(c)))
	pad := r*c - len(normalized)
	for i := 0; i < pad; i++ {
		normalized = append(normalized, ' ')
	}
	transposed := ""
	for i := 0; i < c; i++ {
		line := make([]rune, r)
		for j := 0; j < r; j++ {
			line[j] = normalized[j*c+i]
		}
		if i != 0 {
			transposed += " "
		}
		transposed += string(line)
	}
	return transposed
}
