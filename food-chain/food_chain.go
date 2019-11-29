// Package foodchain implements a solution of the exercise titled `Food Chain'.
package foodchain

import (
	"fmt"
	"strings"
)

var object []string
var secondLine []string

func init() {
	object = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
	secondLine = []string{"",
		"It wriggled and jiggled and tickled inside her.",
		"How absurd to swallow a bird!",
		"Imagine that, to swallow a cat!",
		"What a hog, to swallow a dog!",
		"Just opened her throat and swallowed a goat!",
		"I don't know how she swallowed a cow!",
	}
}

// Verse returns n'th verse of the song.
func Verse(n int) string {
	words := ""
	if n > 0 {
		words += fmt.Sprintf("I know an old lady who swallowed a %s.\n", object[n-1])
		if 1 < n && n < 8 {
			words += secondLine[n-1] + "\n"
		}
		if n <= 7 {
			for i := 1; i < n; i++ {
				words += fmt.Sprintf("She swallowed the %s to catch the %s", object[n-i], object[n-i-1])
				if n-i == 2 {
					words += " that wriggled and jiggled and tickled inside her.\n"
				} else {
					words += ".\n"
				}
			}
			words += "I don't know why she swallowed the fly. Perhaps she'll die."
		} else if n == 8 {
			words += "She's dead, of course!"
		}
	}
	return words
}

// Verses returns the verses between two parameters specified.
func Verses(start, end int) string {
	lyrics := []string{}
	for i := start; i <= end; i++ {
		lyrics = append(lyrics, Verse(i))
	}
	return strings.Join(lyrics, "\n\n")
}

// Song returns the whole song.
func Song() string {
	return Verses(1, 8)
}
