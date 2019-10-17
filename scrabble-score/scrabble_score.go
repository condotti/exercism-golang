package scrabble

import "strings"

// Score returns a score of the WORD in Scrabble game
func Score(word string) int {
	scoreTable := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	score := 0
	word = strings.ToUpper(word)
	for _, letter := range word {
		for key, value := range scoreTable {
			if strings.Index(key, string(letter)) >= 0 {
				score += value
			}
		}
	}
	return score
}
