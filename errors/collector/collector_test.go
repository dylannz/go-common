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

		Context("when nil errors have been collected", func() {
			It("returns nil", func() {
				c := New()
				c.Collect(nil)
				c.Collect(nil)

				actual := c.Errors()

				Expect(actual).To(BeNil())
			})
		})

		Context("when any errors were collected", func() {
			It("returns a Errors error with the collected errors", func() {
				c := New()
				c.Collect(err1)
				c.Collect(err2)

				expected := []error{
					err1,
					err2,
				}

				Expect(c.Errors()).To(Equal(expected))
			})
		})
	})

	Describe("#String (fmt.Stringer)", func() {
		Context("when no errors have been collected", func() {
			It("returns an empty string", func() {
				c := New()

				actual := fmt.Sprint(c)
				expected := ""

				Expect(actual).To(Equal(expected))
			})
		})

		Context("when any errors were collected", func() {
			It("returns the errors concatenated with `; `", func() {
				c := New()
				c.Collect(err1)
				c.Collect(err2)

				actual := fmt.Sprint(c)
				expected := fmt.Sprintf("%s; %s", err1, err2)

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
