package connect

import (
	"fmt"
)

var winner string

func try(stone string, r, c int, rb [][]rune, history map[string]int) {
	// constraints
	state := fmt.Sprintf("%v,%v", r, c)
	if _, ok := history[state]; ok {
		return
	}
	if r < 0 || r >= len(rb) || c < 0 || c >= len(rb[0]) {
		return
	}
	if rb[r][c] != []rune(stone)[0] {
		return
	}
	// goal?
	if stone == "O" && r == len(rb)-1 || stone == "X" && c == len(rb[0])-1 {
		winner = stone
		return
	}
	// step forward to 6 adjacents
	history[state] = 1
	try(stone, r-1, c, rb, history)
	try(stone, r-1, c+1, rb, history)
	try(stone, r, c-1, rb, history)
	try(stone, r, c+1, rb, history)
	try(stone, r+1, c-1, rb, history)
	try(stone, r+1, c, rb, history)
	return
}

func ResultOf(board []string) (string, error) {
	winner = ""

	rb := [][]rune{}
	for _, row := range board {
		rb = append(rb, []rune(row))
	}
	history := map[string]int{}

	for c := range rb[0] {
		try("O", 0, c, rb, history)
	}
	for r := range rb {
		try("X", r, 0, rb, history)
	}
	return winner, nil
}

/*
BenchmarkResultOf-4   	    6996	    186893 ns/op	   26007 B/op	    1852 allocs/op
*/
