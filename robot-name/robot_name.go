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

var names = map[string]bool{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}

// Name assigns a name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(names) >= maxRobots {
			return "", errors.New("name space exhausted")
		}
		r.name = newName()
		for names[r.name] {
			r.name = newName()
		}
		names[r.name] = true
	}
	return r.name, nil
}

// Reset resets the name of robot
func (r *Robot) Reset() {
	r.name = ""
}
