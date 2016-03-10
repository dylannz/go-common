package collector

import (
	"bytes"
	"fmt"
	"sync"
)

// Collector maintains a slice of errors.
type Collector struct {
	lock *sync.RWMutex
	errs []error
}

// New initializes a new ErrorCollector.
func New() Collector {
	return Collector{
		lock: &sync.RWMutex{},
		errs: []error{},
	}
}

// Collect appends a new error to the errors slice.
func (c *Collector) Collect(err error) {
	if err == nil {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.errs = append(c.errs, err)
}

// Errors returns the errors slice or nil if it is empty.
func (c Collector) Errors() []error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if len(c.errs) != 0 {
		return c.errs
	}

	return nil
}

// Error satisfies the error interface by returning an error string, which is the concatenation of all errors
func (c Collector) String() string {
	// Return early with an empty string if we have no errors.
	if len(c.errs) == 0 {
		return ""
	}

	buf := new(bytes.Buffer)
	l := len(c.errs)
	for i, err := range c.errs {
		fmt.Fprintf(buf, "%s", err)
		if (i + 1) < l {
			fmt.Fprint(buf, "; ")
		}
	}
	return buf.String()
}
