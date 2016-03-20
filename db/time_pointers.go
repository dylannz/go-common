package db

import "time"

// NowPointer is a convenience func that returns a pointer to time.Now()
func NowPointer() *time.Time {
	t := time.Now()
	return &t
}
