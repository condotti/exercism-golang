// Package pangram implements a solution for the exercise titled `Pangram'.
package pangram

// IsPangram determines a given sentence is a pangram.
func IsPangram(sentence string) bool {
	seen, count := [26]bool{}, 0
	for _, letter := range sentence {
		switch {
		case 'a' <= letter && letter <= 'z':
			if !seen[letter-'a'] {
				seen[letter-'a'] = true
				count++
			}
		case 'A' <= letter && letter <= 'Z':
			if !seen[letter-'A'] {
				seen[letter-'A'] = true
				count++
			}
		}
		if count >= 26 {
			return true
		}
	}
	return false
}
