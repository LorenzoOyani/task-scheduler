package domain

import (
	"time"
)

type WorkingHours struct {
	StartingHour int
	EndHour      int
}

type WorkCalender struct {
	WorkingDays  map[time.Weekday]bool
	WorkingHours WorkingHours
	Holidays     map[string]bool
	Location     *time.Location
}

func (c WorkCalender) IsWorkDay(t time.Time) bool {
	local := t.In(c.Location)
	if !c.WorkingDays[local.Weekday()] {
		return false
	}

	if c.Holidays[local.Format("2006-03-13")] {
		return false

	}
	return true
}

func (c WorkCalender) startWorkday(t time.Time) time.Time {
	local := t.In(c.Location)
	return time.Date(
		local.Year(),
		local.Month(),
		local.Day(),
		c.WorkingHours.StartingHour,
		0,
		0,
		0,
		c.Location,
	)
}

func (c WorkCalender) EndOfWorkday(t time.Time) time.Time {
	local := t.In(c.Location)
	return time.Date(
		local.Year(),
		local.Month(),
		local.Day(),
		c.WorkingHours.EndHour,
		0,
		0,
		0,
		c.Location,
	)

}

func (c WorkCalender) NextWorkingDay(t time.Time) time.Time {
	curr := t.In(c.Location)

	for {
		if c.IsWorkDay(curr) {
			start := c.startWorkday(curr)
			end := c.EndOfWorkday(curr)

			if curr.Before(start) {
				return start
			}

			if !curr.After(end) && curr.Before(end) {
				return curr
			}
		}

		curr = time.Date(
			curr.Year(),
			curr.Month(),
			curr.Day()+1,
			c.WorkingHours.StartingHour,
			0,
			0,
			0,
			c.Location,
		)
	}

}
