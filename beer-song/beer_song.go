// Package beer implements a solution of the exercise titled `Beer Song'.
package beer

import (
	"errors"
	"fmt"
)

// Verse returns specified verse of the song.
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("verse out of range")
	}

	beer := func(n int) string {
		switch n {
		case 0:
			return "no more bottles"
		case 1:
			return "1 bottle"
		default:
			return fmt.Sprintf("%d bottles", n)
		}
	}
	title := func(s string) string {
		r := []rune(s)
		if r[0] >= 'a' && r[0] <= 'z' {
			r[0] = r[0] - 'a' + 'A'
		}
		return string(r)
	}

	var secondLine string
	if n == 0 {
		secondLine = "Go to the store and buy some more, 99 bottles of beer on the wall."
	} else {
		one := "one"
		if n == 1 {
			one = "it"
		}
		secondLine = fmt.Sprintf("Take %s down and pass it around, %s of beer on the wall.", one, beer(n-1))
	}
	return title(fmt.Sprintf("%s of beer on the wall, %s of beer.\n%s\n", beer(n), beer(n), secondLine)), nil
}

// Verses return the verses in the range.
func Verses(upperBound, lowerBound int) (string, error) {
	if upperBound < lowerBound {
		return "", errors.New("upper bound < lower bound")
	}
	verses := ""
	for n := upperBound; n >= lowerBound; n-- {
		verse, err := Verse(n)
		if err != nil {
			return "", err
		}
		verses += verse + "\n"
	}
	return verses, nil
}

// Song returns a whole song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}

/*
BenchmarkSeveralVerses-4   	    2794	    437702 ns/op	  740736 B/op	    1682 allocs/op
BenchmarkEntireSong-4      	    2864	    446979 ns/op	  733537 B/op	    1583 allocs/op
*/
