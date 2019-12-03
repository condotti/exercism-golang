// Package allergies implements a solution of the exercise titled `Allergies'.
package allergies

var allergySource []string

func init() {
	allergySource = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes",
		"chocolate", "pollen", "cats"}
}

// Allergies returns allergy sources for given score.
func Allergies(score uint) []string {
	sources := []string{}
	for _, item := range allergySource {
		if score&1 == 1 {
			sources = append(sources, item)
		}
		score >>= 1
	}
	return sources
}

// AllergicTo receives a score and a source, determines if the source causes allergie.
func AllergicTo(score uint, source string) bool {
	for _, item := range Allergies(score) {
		if item == source {
			return true
		}
	}
	return false
}
