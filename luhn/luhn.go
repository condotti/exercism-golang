// Package luhn implements a solution of the exercise titled `Luhn'.
package luhn

// Valid determines whether or not NUM is valid per the Luhn formula
func Valid(num string) bool {
	doubled := []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	rnum := []rune(num)
	var i, sum int
	for j := range rnum {
		c := rnum[len(rnum)-j-1]
		if c != ' ' {
			n := int(c - '0')
			if 0 <= n && n <= 9 {
				if i&1 == 0 {
					sum += n
				} else {
					sum += doubled[n]
				}
				i++
			} else {
				return false
			}
		}
	}
	if i <= 1 {
		return false
	}
	return sum%10 == 0
}
