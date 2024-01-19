package sslcertchecker

import "time"

func calcDaysDiff(d time.Time, u time.Time) int64 {
	return int64(d.Sub(u).Hours() / 24)
}
