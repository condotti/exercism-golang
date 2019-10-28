// Package scrabble implements a solution of exercise titled `Scrabble Score'.
package scrabble

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
	var scoreVector [26]int

	for letters, score := range scoreTable {
		for _, letter := range letters {
			scoreVector[letter-'A'] = score
		}

	}

	total := 0
	for _, letter := range word {
		if 'A' <= letter && letter <= 'Z' {
			total += scoreVector[letter-'A']
		} else if 'a' <= letter && letter <= 'z' {
			total += scoreVector[letter-'a']
		}
	}
	return total
}
