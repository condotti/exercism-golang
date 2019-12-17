package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a robot
type Robot struct {
	name string
}

const maxRobots = 26 * 26 * 1000

var names map[string]int

func init() {
	rand.Seed(time.Now().UnixNano())
	names = map[string]int{}
}

// String is the stringer method.
func (r Robot) String() string {
	return r.name
}

// Name assigns a name.
func (r *Robot) Name() (string, error) {
	name := func(n int) string {
		return fmt.Sprintf("%c%c%03d", n/26000+'A', n/1000%26+'A', n%1000)
	}

	if r.name == "" {
		if len(names) >= maxRobots {
			return "", errors.New("name space exhausted")
		}
		var n string
		for n = name(rand.Intn(maxRobots)); names[n] != 0; n = name(rand.Intn(maxRobots)) {
		}
		names[n] = 1
		r.name = n
	}
	return r.String(), nil
}

// Reset resets the name of robot
func (r *Robot) Reset() {
	delete(names, r.name)
	// Deleting the entry of the directory is required or not?
	// Name conflict was reported when deleted on the benchmark.
	// Deleting the entry causes name space depletion on the benchmark.
	r.name = ""
	r.Name()
}
