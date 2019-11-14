// Package secret implements a solution of the exercise titled `Secret Handshake'.
package secret

// Handshake converts a number into the appropriate sequence of events for a secret handshake.
func Handshake(n uint) []string {
	actions := []string{"wink", "double blink", "close your eyes", "jump"}
	hs := []string{}
	for i, action := range actions {
		if n>>i&1 == 1 {
			hs = append(hs, action)
		}
	}
	if n>>len(actions)&1 == 1 {
		hs = func(s []string) []string {
			r := []string{}
			for _, a := range s {
				r = append([]string{a}, r...)
			}
			return r
		}(hs)
	}
	return hs
}
