package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a robot
type Robot int

const maxRobots = 26 * 26 * 1000

var conflict [maxRobots]bool

var nRobot int

func init() {
	rand.Seed(time.Now().UnixNano())
}

// String is the stringer method.
func (r *Robot) String() string {
	n := ^int(*r)
	return fmt.Sprintf("%c%c%03d", n/26000%26+'A', n/1000%26+'A', n%1000)
}

// Name assigns a name.
func (r *Robot) Name() (string, error) {
	if *r == 0 {
		if nRobot < maxRobots {
			var n int
			for n = rand.Intn(maxRobots); conflict[n]; n = rand.Intn(maxRobots) {}
			conflict[n] = true
			*r = Robot(^n)
			nRobot++
		} else {
			return "", errors.New("name exhausted")
		}
	}
	return r.String(), nil
}

// Reset resets the name of robot
func (r *Robot) Reset() {
	conflict[^int(*r)] = false
	nRobot--
	*r = 0
	r.Name()
}
