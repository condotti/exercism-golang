package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	gigasecond, _ := time.ParseDuration("1000000000s")
	return t.Add(gigasecond)
}
