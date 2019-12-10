// Package railfence implements a solution of exercise titled `Rail Fence Cipher'.
package railfence

import "strings"

func cycle(rails int) []int {
	cycle := []int{}
	for i := range make([]int, rails-1) {
		cycle = append(cycle, i)
	}
	for i := range make([]int, rails-1) {
		cycle = append(cycle, rails-i-1)
	}
	return cycle
}

// Encode encodes a messeage with specified number of rails.
func Encode(message string, rails int) string {
	cycle := cycle(rails)
	fence := make([]string, rails)
	for i, c := range message {
		fence[cycle[i%len(cycle)]] += string(c)
	}
	return strings.Join(fence, "")
}

// Decode decodes a message with specified number of rails.
func Decode(message string, rails int) string {
	cycle := cycle(rails)
	fence := make([][]int, rails)
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
