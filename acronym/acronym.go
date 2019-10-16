package acronym

import (
	"strings"
	"regexp"
)

func Abbreviate(s string) string {
	return strings.ToUpper(regexp.MustCompile(`([A-Za-z])(\w|')*(\W|_)*`).ReplaceAllString(s, "$1"))
}
