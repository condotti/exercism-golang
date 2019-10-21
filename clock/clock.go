package clock

import "fmt"

type Clock int

// New(H, M) is the constuctor of Clock which represents H:M
func New(h, m int) Clock {
	m += h*60
	m %= 60*24
	if m < 0 {
		m += 60*24
	}
	return Clock(m);
}

// String() is a stringer for Clock
func (c Clock)String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add(M) returns M minutes added to the clock instance
func (c Clock)Add(m int) Clock {
	return New(0, int(c) + m)
}

// Subtract(M) returns M minutes subtracted from the clock instance
func (c Clock)Subtract(m int) Clock {
	return New(0, int(c) - m)
}
