// Package wordcount implements a solution for the exercise titled `Word Count'.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map of word-occurence association.
type Frequency map[string]int

// WordCount counts the occurrences of each word in that phrase.
func WordCount(s string) Frequency {
	freq := Frequency{}
	re := regexp.MustCompile(`([a-zA-Z0-9][a-zA-Z0-9']*[a-zA-Z0-9]|[a-zA-Z0-9]+)`)
	for _, word := range re.FindAllString(s, -1) {
		freq[strings.ToLower(word)]++
	}
	return freq
}
