// Package encode implements a solution for the exercise titiled `Run Length Encoding'.
package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// RunLengthEncode returns encoded string of the input
func RunLengthEncode(s string) string {
	type pair struct {
		c rune
		n int
	}
	pairs := []pair{}
	for _, c := range s {
		if len(pairs) > 0 && pairs[len(pairs)-1].c == c {
			pairs[len(pairs)-1].n++
		} else {
			pairs = append(pairs, pair{c, 1})
		}
	}
	encoded := ""
	for _, p := range pairs {
		if p.n == 1 {
			encoded += string(p.c)
		} else {
			encoded += fmt.Sprintf("%d%c", p.n, p.c)
		}
	}
	return encoded
}

// RunLengthDecode returns decoded string of the input
func RunLengthDecode(s string) string {
	re := regexp.MustCompile(`([0-9]*)([a-zA-Z ])`)
	decoded := ""
	for _, sub := range re.FindAllStringSubmatch(s, -1) {
		if i, err := strconv.Atoi(sub[1]); err == nil {
			decoded += strings.Repeat(sub[2], i)
		} else {
			decoded += sub[2]
		}
	}
	return decoded
}
