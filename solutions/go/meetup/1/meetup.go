package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Last   WeekSchedule = "last"
	Teenth WeekSchedule = "teenth"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	fisrtdayofmonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()
	switch wSched {
	case "first":
		{
			return offsetForWeekday(wDay, fisrtdayofmonth, 1)
		}
	case "second":
		{
			return offsetForWeekday(wDay, fisrtdayofmonth, 2)
		}
	case "third":
		{
			return offsetForWeekday(wDay, fisrtdayofmonth, 3)
		}
	case "fourth":
		{
			return offsetForWeekday(wDay, fisrtdayofmonth, 4)
		}
	case "last":
		{
			daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
			lastdayofmonth := time.Date(year, month, daysInMonth, 0, 0, 0, 0, time.UTC).Weekday()
			offset := (int(lastdayofmonth) - int(wDay) + 7) % 7
			return daysInMonth - offset
		}
	case "teenth":
		{
			for day := 13; day <= 19; day++ {
				if time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Weekday() == wDay {
					return day
				}
			}
		}

	}
	return 0
}

func offsetForWeekday(wDay, fisrtdayofmonth time.Weekday, n int) int {
	offset := (int(wDay) - int(fisrtdayofmonth) + 7) % 7
	return offset + (n-1)*7+1
}
