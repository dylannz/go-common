package collector

import (
	"bytes"
	"fmt"
	"sync"
)

// Collector maintains a slice of errors.
type Collector struct {
	lock *sync.RWMutex
	errs Errors
}

// New initializes a new ErrorCollector.
func New() Collector {
	return Collector{
		lock: &sync.RWMutex{},
		errs: Errors{
			Errors: []error{},
		},
	}
}

// Collect appends a new error to the errors slice.
func (c *Collector) Collect(err error) {
	if err == nil {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.errs.append(err)
}

// Errors returns the errors slice or nil if it is empty.
func (c Collector) Errors() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.errs.len() != 0 {
		return c.errs
	}

	return nil
}

func (c Collector) String() string {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.errs.Error()
}

// Errors is slice of errors that satisfies the error interface.
type Errors struct {
	Errors []error
}

// Error satisfies the error interface by returning an error string, which is the concatenation of all errors
func (e Errors) Error() string {
	buf := new(bytes.Buffer)
	l := e.len()
	for i, v := range e.list() {
		fmt.Fprintf(buf, "%s", v)
		if (i + 1) < l {
			fmt.Fprint(buf, "; ")
		}
	}
	return buf.String()
}

// len returns the count of errors.
func (e Errors) len() int {
	return len(e.list())
}

// Llist returns the error slice
func (e Errors) list() []error {
	return e.Errors
}

// append adds a new error to the end of the error slice.
func (e *Errors) append(err error) {
	e.Errors = append(e.Errors, err)
}
