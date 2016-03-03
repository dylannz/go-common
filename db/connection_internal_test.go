package db

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PG", func() {
	Describe("logSafeConnectionString", func() {
		os.Setenv("DB_NAME", "test")
		os.Setenv("DB_USER", "user")
		os.Setenv("DB_HOST", "host")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_PASSWORD", "password")
		os.Setenv("DB_SSL_MODE", "disable")

		Context("when there is a password", func() {
			It("replaces the password with stars", func() {
				db := PG{}

				actual := db.logSafeConnectionString()
				expected := "postgres://user:****@host:5432/test?sslmode=disable"

				Expect(actual).To(Equal(expected))
			})
		})

		Context("when there is no password", func() {
			It("leaves the connection string as is", func() {
				os.Setenv("DB_PASSWORD", "")

				db := PG{}

				actual := db.logSafeConnectionString()
				expected := db.connectionString()

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
