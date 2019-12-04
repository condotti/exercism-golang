// Package piglatin implements a solution of the exercise titled `Pig Latin'.
package piglatin

import (
	"regexp"
	"strings"
)

// Sentence translates from English to Pig Latin.
func Sentence(source string) string {
	pattern := []string{
		`^([aeiou].*|yt.*|xr.*)()$`, // Rule 1
		`^([^aeiouq]*qu)(.*)$`,      // Rule 3
		`^([^aeiouy]+)(y.*)$`,       // Rule 4
		`^([^aeiou]+)(.*)$`,         // Rule 2
	}
	rule := []*regexp.Regexp{}
	for _, pat := range pattern {
		rule = append(rule, regexp.MustCompile(pat))
	}
	translated := []string{}
	for _, word := range strings.Split(source, " ") {
		translated = append(translated, func(word string) string {
			for _, rex := range rule {
				if rex.MatchString(word) {
					return rex.ReplaceAllString(word, "$2$1") + "ay"
				}
			}
			return "" // should not reach here
		}(word))
	}
	return strings.Join(translated, " ")
}
