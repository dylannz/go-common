package db_test

import (
	"github.com/HomesNZ/go-common/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PG Array Test", func() {
	Context("converting row value", func() {
		It("returns the correct number of elements", func() {
			value := "{string_1, string_2, string_3 }"

			actual := db.ParseArray(value)
			expected := []string{"string_1", "string_2", "string_3"}

			Expect(actual).To(Equal(expected))
		})
	})
	Context("converting a splice/arry to string for Postgres", func() {
		It("is a string", func() {

			value := []string{"string_1", "string_2", "string_3"}

			actual := db.CreateStringArray(value)
			expected := `{"string_1", "string_2", "string_3"}`

			Expect(actual).To(Equal(expected))
		})
	})
})
