// Package isbn implements a solution of the exercise titled `ISBN Verifier'
package isbn

// IsValidISBN determines if given digits are valid or not.
func IsValidISBN(digits string) bool {
	i, sum := 10, 0
	for _, d := range digits {
		switch {
		case i < 1:
			return false // digits too long
		case '0' <= d && d <= '9':
			sum += int(d - '0') * i
			i--
		case i == 1 && d == 'X':
			sum += 10
			i--
		case d != '-':
			return false // invalid char
		}
	}
	if i == 0 {
		return sum%11 == 0
	}
	return false // digits too short
}
