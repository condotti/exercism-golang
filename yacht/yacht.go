// Package yacht implements a solution of the exercise titled `Yacht'.
package yacht

import "sort"

// Score returns the score of given dices and a category.
func Score(dice []int, category string) int {
	score := 0
	numbers := []string{"ones", "twos", "threes", "fours", "fives", "sixes"}
	t := 0
	for i, c := range numbers {
		if c == category {
			t = i + 1
			break
		}
	}
	if t > 0 {
		for _, d := range dice {
			if d == t {
				score += d
			}
		}
		return score
	}
	sort.Slice(dice, func(i, j int) bool { return dice[i] < dice[j] })
	switch category {
	case "full house":
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[3] == dice[4] && dice[2] != dice[3] {
			return dice[0]*3 + dice[3]*2
		} else if dice[0] == dice[1] && dice[2] == dice[3] && dice[3] == dice[4] && dice[1] != dice[2] {
			return dice[0]*2 + dice[2]*3
		} else {
			return 0
		}
	case "four of a kind":
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] {
			return dice[0] * 4
		} else if dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4] {
			return dice[1] * 4
		} else {
			return 0
		}
	case "little straight":
		for i, d := range dice {
			if d != i+1 {
				return 0
			}
		}
		return 30
	case "big straight":
		for i, d := range dice {
			if d != i+2 {
				return 0
			}
		}
		return 30
	case "choice":
		for _, d := range dice {
			score += d
		}
		return score
	case "yacht":
		for _, d := range dice[1:] {
			if d != dice[0] {
				return 0
			}
		}
		return 50
	}
	return 0
}
