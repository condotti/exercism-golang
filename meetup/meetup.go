// Package meetup implements a solution of the exercise titled `Meetup'.
package meetup

import "time"

// WeekSchedule represents the date descriptor.
type WeekSchedule int

// Names of week
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day calculate the date of meetups.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	firstDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	daysAfter := (int(weekday) - int(firstDate.Weekday()) + 7) % 7
	targetDates := []time.Time{}
	for date := firstDate.AddDate(0, 0, daysAfter); date.Month() == month; date = date.AddDate(0, 0, 7) {
		targetDates = append(targetDates, date)
	}
	switch week {
	case Last:
		return targetDates[len(targetDates)-1].Day()
	case Teenth:
		for i := range targetDates {
			if targetDates[len(targetDates)-i-1].Day() < 20 {
				return targetDates[len(targetDates)-i-1].Day()
			}
		}
	default:
		return targetDates[week].Day()
	}
	return 0 // should not reach here
}
