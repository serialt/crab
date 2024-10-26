package crab

import "time"

const (
	CrabDateTimeV2 = "2006/01/02 15:04:05"
	CrabSecond     = "20060102150405"
	CrabDay        = "20060102"
)

// GetCrabSecond 20060102150405
func GetCrabSecond(t time.Time) string {
	return t.Format(CrabSecond)
}

// GetCrabDay 20060102
func GetCrabDay(t time.Time) string {
	return t.Format(CrabDay)
}

// GetDateTime 2006-01-02 15:04:05
func GetDateTime(t time.Time) string {
	return t.Format(time.DateTime)
}

// GetDateTimeV2 2006/01/02 15:04:05
func GetDateTimeV2(t time.Time) string {
	return t.Format(CrabDateTimeV2)
}
