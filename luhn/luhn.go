package luhn

// Valid determines whether or not NUM is valid per the Luhn formula
func Valid(num string) bool {
	digits := []int{}
	for _, c := range num {
		if c != ' ' {
			if c >= '0' && c <= '9' {
				digits = append(digits, int(c-'0'))
			} else {
				return false
			}
		}
	}
	if len(digits) <= 1 {
		return false
	}
	sum := 0
	doubled := []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	for i := range digits {
		if i&1 == 0 {
			sum += digits[len(digits)-i-1]
		} else {
			sum += doubled[digits[len(digits)-i-1]]
		}
	}
	return sum%10 == 0
}
