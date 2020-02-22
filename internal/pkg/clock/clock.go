package clock

import (
	"github.com/twitchtv/twirp"
	"time"
)

// 9999-12-31 00:00:00.000 (MySQL max value for Date)
func InfinityAt() time.Time {
	return time.Date(9999, 12, 31, 0,0,0,0,time.UTC)
}

// 1970-01-01 00:00:00.000
func ZeroAt() time.Time {
	return time.Unix(0, 0)
}

func ToMonthDate(month time.Month, day int) (time.Time, error) {
	if day <= 0 || day > 31 {
		return time.Time{}, twirp.InvalidArgumentError("day", "invalid")
	}
	if month <= 0 || month > 12 {
		return time.Time{}, twirp.InvalidArgumentError("month", "invalid")
	}
	current := time.Now().UTC()
	return time.Date(current.Year(), month, day, 0, 0,0,0, time.UTC), nil
}
