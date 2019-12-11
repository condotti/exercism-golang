// Package diamond implements a solution of the exercise titled `Diamond'.
package diamond

import (
	"errors"
	"strings"
)

// Gen returns a slice of strings representing required diamond diagram.
func Gen(c byte) (string, error) {
	if c < 'A' || c > 'Z' {
		return "", errors.New("invalid argument")
	}
	diamond := []string{"A"}
	size := int(c - 'A')
	for i := 1; i <= size; i++ {
		tmp := []string{}
		for _, l := range diamond {
			tmp = append(tmp, " "+l+" ")
		}
		tmp = append(tmp, string('A'+i)+strings.Repeat(" ", i*2-1)+string('A'+i))
		diamond = tmp
	}
	for i := 0; i < size; i++ {
		diamond = append(diamond, diamond[size-i-1])
	}
	return strings.Join(diamond, "\n") + "\n", nil
}
