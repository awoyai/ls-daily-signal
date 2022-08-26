package utils

import "time"

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006/01/02"
)

var loc, _ = time.LoadLocation("Local")

func StrToDate(str string) (time.Time, error) {
	return time.ParseInLocation(DateLayout, str, loc)
}

func StrToTime(str string) (time.Time, error) {
	return time.ParseInLocation(TimeLayout, str, loc)
}
