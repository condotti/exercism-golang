package bob

import "regexp"

func Hey(remark string) string {
	rem := []byte(remark)
	re := []*regexp.Regexp {
		regexp.MustCompile(`^\s*$`),
		regexp.MustCompile(`[A-Z]`),
		regexp.MustCompile(`[a-z]`),
		regexp.MustCompile(`\?\s*$`),
	}
	if re[0].Match(rem) {
		return "Fine. Be that way!"
	} else if re[1].Match(rem) && !re[2].Match(rem) {
		if (re[3].Match(rem)) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else if re[3].Match(rem) {
		return "Sure."
	} else {
		return "Whatever."
	}
}
