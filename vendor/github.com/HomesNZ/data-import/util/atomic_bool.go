package util

import "sync/atomic"

// AtomicBool is a concurrent safe bool implementation that atomically updates an int32 to either a 1 or 0.
type AtomicBool struct{ flag int32 }

// Set updates the value for the AtomicBool
func (a *AtomicBool) Set(value bool) {
	var i int32
	if value {
		i = 1
	}
	atomic.StoreInt32(&(a.flag), int32(i))
}

// Get returns the value of the AtomicBool
func (a *AtomicBool) Get() bool {
	if atomic.LoadInt32(&(a.flag)) != 0 {
		return true
	}
	return false
}
