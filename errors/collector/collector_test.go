package collector_test

import (
	"errors"
	"fmt"

	. "github.com/HomesNZ/go-common/errors/collector"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collector", func() {
	err1 := errors.New("Error 1")
	err2 := errors.New("Error 2")

	Describe("#Errors", func() {
		Context("when no errors have been collected", func() {
			It("returns nil", func() {
				c := New()

				Expect(c.Errors()).To(BeNil())
			})
		})

		Context("when any errors were collected", func() {
			It("returns nil", func() {
				c := New()
				c.Collect(err1)
				c.Collect(err2)

				expected := []error{
					err1,
					err2,
				}

				actual := c.Errors().(Errors)

				Expect(actual.Errors).To(Equal(expected))
			})
		})
	})

	Describe("#String", func() {
		Context("when no errors have been collected", func() {
			It("returns `<nil>`", func() {
				c := New()

				actual := fmt.Sprint(c.Errors())
				expected := "<nil>"

				Expect(actual).To(Equal(expected))
			})
		})

		Context("when any errors were collected", func() {
			It("returns the errors concatenated with `; `", func() {
				c := New()
				c.Collect(err1)
				c.Collect(err2)

				actual := fmt.Sprint(c.Errors())
				expected := fmt.Sprintf("%s; %s", err1, err2)

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
