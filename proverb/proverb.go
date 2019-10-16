package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	proverb := []string{}
	if len(rhyme) > 0 {
		for i := range rhyme {
			if i > 0 {
				proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
			}
		}
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return proverb
}
