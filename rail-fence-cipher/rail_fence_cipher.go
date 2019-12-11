// Package railfence implements a solution of exercise titled `Rail Fence Cipher'.
package railfence

import "strings"

func cycle(n int) (cycle []int) {
	for i := 0; i < n-1; i++ {
		cycle = append(cycle, i)
	}
	for i := 0; i < n-1; i++ {
		cycle = append(cycle, n-i-1)
	}
	return
}

// Encode encodes a messeage with specified number of rails.
func Encode(message string, rails int) string {
	fence := make([]string, rails)
	cycle := cycle(rails)
	for i, c := range message {
		fence[cycle[i%len(cycle)]] += string(c)
	}
	return strings.Join(fence, "")
}

// Decode decodes a message with specified number of rails.
func Decode(message string, rails int) string {
	fence := make([][]int, rails)
	cycle := cycle(rails)
	for i := range message {
		fence[cycle[i%len(cycle)]] = append(fence[cycle[i%len(cycle)]], i)
	}
	pos := []int{}
	for _, line := range fence {
		pos = append(pos, line...)
	}
	decoded := make([]rune, len(message))
	for i, c := range message {
		decoded[pos[i]] = c
	}
	return string(decoded)
}
