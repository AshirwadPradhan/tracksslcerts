package helpers

import "time"

func CalcDaysDiff(d time.Time, u time.Time) int64 {
	return int64(d.Sub(u).Hours() / 24)
}
