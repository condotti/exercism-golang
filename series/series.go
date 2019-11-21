// Package series implements a solution of the exercise titled `Series'.
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	subs := []string{}
	if n <= len(s) {
		for i := 0; i < len(s)-n+1; i++ {
			subs = append(subs, s[i:i+n])
		}
	}
	return subs
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

// First returns the firsst substring of s with length n, or signals if s is shorter than n.
func First(n int, s string) (string, bool) {
	if len(s) >= n {
		return UnsafeFirst(n, s), true
	}
	return "", false
}
