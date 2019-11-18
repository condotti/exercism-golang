// Package anagram implements a solution for the exercise titled `Anagram'.
package anagram

import (
	"sort"
	"strings"
)

// Detect returns a list of words containing anagrams of subject from given candidates.
func Detect(subject string, candidates []string) []string {
	signature := func(s string) string {
		a := []byte(s)
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		return string(a)
	}
	lowerSubj := strings.ToLower(subject)
	sigSubj := signature(lowerSubj)
	anagrams := []string{}
	for _, word := range candidates {
		lowerWord := strings.ToLower(word)
		if lowerSubj != lowerWord && sigSubj == signature(lowerWord) {
			anagrams = append(anagrams, word)
		}
	}
	return anagrams
}
