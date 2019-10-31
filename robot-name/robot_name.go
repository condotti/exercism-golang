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

func search(min, max int) int {
	if min == max {
		if conflict[min] {
			return -1 // not found
		}
		return min // found
	}
	n := (min + max) / 2
	if conflict[n] {
		nn := search(min, n)
		if nn == -1 {
			return search(n, max)
		}
		return nn // found
	}
	return n // found
}

// String is the stringer method.
func (r *Robot) String() string {
	n := ^int(*r)
	return fmt.Sprintf("%c%c%03d", n/26000%26+'A', n/1000%26+'A', n%1000)
}

// Name assigns a name.
func (r *Robot) Name() (string, error) {
	if *r == 0 {
		var n int
		if nRobot < maxRobots/2 {
			for n = rand.Intn(maxRobots); conflict[n]; n = rand.Intn(maxRobots) {
			}
		} else if nRobot < maxRobots {
			n = search(0, maxRobots)
		} else {
			return "", errors.New("name exhausted")
		}
		conflict[n] = true
		*r = Robot(^n)
		nRobot++
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
