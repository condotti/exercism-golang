// Package brackets implements a solution of the exercise titled `Matching Brackets'.
package brackets

// Bracket determines if brackets, braces, parentheses in the string is balanced or not.
func Bracket(input string) bool {
	pairs := map[rune]rune{'}': '{', ']': '[', ')': '('}
	stack := []rune{}
	for _, c := range input {
		switch c {
		case '{', '[', '(':
			stack = append(stack, c)
		case '}', ']', ')':
			if len(stack) > 0 && stack[len(stack)-1] == pairs[c] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
