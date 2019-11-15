package isogram

// IsIsogram determines if a WORD or phrase is an isogram
func IsIsogram(word string) bool {
	var seen [26]bool
	for _, letter := range word {
		if 'a' <= letter && letter <= 'z' {
			if seen[letter-'a'] {
				return false
			}
			seen[letter-'a'] = true
		} else if 'A' <= letter && letter <= 'Z' {
			if seen[letter-'A'] {
				return false
			}
			seen[letter-'A'] = true
		} else if letter != ' ' && letter != '-' {
			return false
		}
	}
	return true
}
