package helper

import "time"

// InTimeSpan will check whether cancelation time passed or not
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
