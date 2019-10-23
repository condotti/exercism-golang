// Package etl implements a solution of the exercise titled `ETL'.
package etl

import "strings"

// Transform returns new data representation of given LEGACY data.
func Transform(legacy map[int][]string) map[string]int {
	fresh := map[string]int{}
	for score, letters := range legacy {
		for _, letter := range letters {
			fresh[strings.ToLower(letter)] = score
		}
	}
	return fresh
}
