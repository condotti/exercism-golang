// Package atbash implements a solution of the exercise titled `Atbash Cipher'.
package atbash

// Atbash returns encoded string of the input
func Atbash(plain string) string {
	encoded := []rune{}
	var n rune
	for _, c := range plain {
		n = 0
		switch {
		case 'a' <= c && c <= 'z':
			n = 'a' + 'z' - c
		case 'A' <= c && c <= 'Z':
			n = 'a' + 'Z' - c
		case '0' <= c && c <= '9':
			n = c
		}
		if n > 0 {
			if len(encoded)%6 == 5 {
				encoded = append(encoded, ' ')
			}
			encoded = append(encoded, n)
		}
	}
	return string(encoded)
}
