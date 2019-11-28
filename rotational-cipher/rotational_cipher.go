// Package rotationalcipher implements a solution of the exercise titled `Rotational Cipher'.
package rotationalcipher

func isLower(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func isUpper(c rune) bool {
	return isLower(c - 'A' + 'a')
}

func nonNegMod(a, b rune) rune {
	return (b + (a % b)) % b
}

// RotationalCipher encodes a plain text with the shift key specified.
func RotationalCipher(plain string, key int) string {
	cipher := []rune{}
	for _, letter := range plain {
		switch {
		case isLower(letter):
			cipher = append(cipher, nonNegMod(letter-'a'+rune(key), 26)+'a')
		case isUpper(letter):
			cipher = append(cipher, nonNegMod(letter-'A'+rune(key), 26)+'A')
		default:
			cipher = append(cipher, letter)
		}
	}
	return string(cipher)
}

// BenchmarkRotationalCipher-4   	  403110	      2558 ns/op	    1272 B/op	      34 allocs/op
