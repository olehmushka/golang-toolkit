package timetools

import "time"

func TimeToString(in time.Time) string {
	return in.Format(DefaultTimeFormat)
}

func StringToTime(in string) (time.Time, error) {
	return time.Parse(DefaultTimeFormat, in)
}
