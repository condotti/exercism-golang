// Package clock implements a solution of the exercise titled `Clock'.
package clock

import "fmt"

// Clock represents a clock that handles times without dates.
type Clock int

// New is the ctor of Clock.
func New(h, m int) Clock {
	m += h * 60
	m %= 60 * 24
	if m < 0 {
		m += 60 * 24
	}
	return Clock(m)
}

// String is a stringer of Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add returns M minutes added to the clock instance
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract returns M minutes subtracted from the clock instance
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}
