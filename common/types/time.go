package types

import "time"

const (
	TimeFormatDate = "2006-01-02 15:04:05"
)

func GetDayStartEn() (start time.Time, end time.Time) {
	now := time.Now()
	start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 50, 999, time.Local)
	return
}

func GetMonthStartEnd() (start time.Time, end time.Time) {
	now := time.Now()
	start = now.AddDate(0, 0, -now.Day()+1)
	end = now.AddDate(0, 1, -now.Day())
	return
}
