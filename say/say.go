// Package say implements a solution of the exercise titled `Say'.
package say

var (
	scale, unit, teenth, tenth []string
)

func init() {
	scale = []string{"", "thousand", "million", "billion"}
	unit = []string{"", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}
	teenth = []string{"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "ninetten"}
	tenth = []string{"", "", "twenty", "thirty", "forty",
		"fifty", "sixty", "seventy", "eighty", "ninety"}
}

// step1
func under99(n int) string {
	switch {
	case n < 20:
		return append(unit, teenth...)[n]
	case n%10 == 0:
		return tenth[n/10]
	}
	return tenth[n/10] + "-" + unit[n%10]
}

func under999(n int) string {
	switch {
	case n < 99:
		return under99(n)
	case n%100 == 0:
		return unit[n/100] + " hundred"
	}
	return unit[n/100] + " hundred " + under99(n%100)
}

// step2
func breakUp(n int64) ([]int, bool) {
	if n < 0 || n >= 1e12 {
		return nil, false
	}
	return []int{int(n / 1e9), int((n / 1e6) % 1e3), int((n / 1e3) % 1e3), int(n % 1e3)}, true
}

// Say spells out a given number in English. (step3 and 4)
func Say(n int64) (string, bool) {
	chunks, ok := breakUp(n)
	if !ok {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	english := ""
	for i, chunk := range chunks {
		if chunk != 0 {
			if len(english) != 0 {
				english += " "
			}
			english += under999(chunk)
			if len(scale[len(scale)-i-1]) != 0 {
				english += " " + scale[len(scale)-i-1]
			}
		}
	}
	return english, true
}

/*
BenchmarkSay-4   	  289518	      4303 ns/op	    3872 B/op	      38 allocs/op
*/
