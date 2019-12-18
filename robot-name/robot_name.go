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

var names = map[string]int{}

// String is the stringer method.
func (r Robot) String() string {
	return r.name
}

// Name assigns a name.
func (r *Robot) Name() (string, error) {
	newName := func() string {
		r1 := string(rand.Intn(26) + 'A')
		r2 := string(rand.Intn(26) + 'A')
		num := rand.Intn(1000)
		return fmt.Sprintf("%s%s%03d", r1, r2, num)
	}

	if r.name == "" {
		rand.Seed(time.Now().UnixNano())
		if len(names) >= maxRobots {
			return "", errors.New("name space exhausted")
		}
		var name string
		for name = newName(); names[name] != 0; name = newName() {
		}
		names[name] = 1
		r.name = name
	}
	return r.String(), nil
}

// Reset resets the name of robot
func (r *Robot) Reset() {
	r.name = ""
}
