// Package wordy implements a solution of the exercise titled `Wordy'.
package wordy

import (
	"regexp"
	"strconv"
)

// Answer parses and evaluates simple math word problems returning the answer as an integer.
func Answer(problem string) (int, bool) {
	match := regexp.MustCompile(`^\D+\s+(-?\d+)(.*)\?$`).FindAllStringSubmatch(problem, -1)
	if len(match) != 1 {
		return 0, false
	}
	answer, _ := strconv.Atoi(match[0][1])
	rest := match[0][2]

	for len(rest) > 0 {
		match = regexp.MustCompile(`^\s*(\D+)\s+(-?\d+)(.*)$`).FindAllStringSubmatch(rest, -1)
		if len(match) != 1 {
			return 0, false
		}
		n, _ := strconv.Atoi(match[0][2])
		rest = match[0][3]
		switch match[0][1] {
		case "plus":
			answer += n
		case "minus":
			answer -= n
		case "multiplied by":
			answer *= n
		case "divided by":
			answer /= n
		default:
			return 0, false
		}
	}
	return answer, true
}

/*
BenchmarkAnswer-4   	    1743	    636279 ns/op	  550627 B/op	    5462 allocs/op
*/
