// Package wordy implements a solution of the exercise titled `Wordy'.
package wordy

import (
	"regexp"
	"strconv"
)

// Answer parses and evaluates simple math word problems returning the answer as an integer.
func Answer(problem string) (int, bool) {
	var err interface{}
	var answer, n int
	var oper string

	reProb := regexp.MustCompile(`^What is (-?\d+)(.*)\?$`)
	reOper := regexp.MustCompile(`^ (plus|minus|multiplied by|divided by) (-?\d+)(.*)$`)

	match := reProb.FindAllStringSubmatch(problem, -1)
	if len(match) != 1 {
		return 0, false
	}
	num, rest := match[0][1], match[0][2]
	if answer, err = strconv.Atoi(num); err != nil {
		return 0, false
	}

	for len(rest) > 0 {
		match = reOper.FindAllStringSubmatch(rest, -1)
		if len(match) != 1 {
			return 0, false
		}
		oper, num, rest = match[0][1], match[0][2], match[0][3]
		if n, err = strconv.Atoi(num); err != nil {
			return 0, false
		}

		switch oper {
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
