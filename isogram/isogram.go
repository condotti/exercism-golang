package isogram

import "strings"

// IsIsogram determines if a WORD or phrase is an isogram
func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	letterMap := make(map[rune]bool)
	for _, letter := range word {
		if strings.Index(" -", string(letter)) < 0 {
			if letterMap[letter] {
				return false
			}
		}
		letterMap[letter] = true
	}
	return true
}
