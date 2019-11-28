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
	cipher := []rune(plain)
	for i, _ := range cipher {
		switch {
		case isLower(cipher[i]):
			cipher[i] = nonNegMod(cipher[i]-'a'+rune(key), 26) + 'a'
		case isUpper(cipher[i]):
			cipher[i] = nonNegMod(cipher[i]-'A'+rune(key), 26) + 'A'
		}
	}
	return string(cipher)
}

// BenchmarkRotationalCipher-4   	  710364	      1775 ns/op	     320 B/op	      10 allocs/op
