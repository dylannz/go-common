package redis_test

import (
	. "github.com/HomesNZ/data-import/cache/redis"
	"github.com/rafaeljusto/redigomock"

	"github.com/garyburd/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redis", func() {
	Describe(".Conn", func() {
		Context("with an external connection [integration]", func() {
			It("returns an active connection to Redis", func() {
				c := CacheConn()

				_, err := c.Conn().Do("PING")

				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("with an mocked connection", func() {
			It("returns an active connection to Redis", func() {
				// TODO (nicday): Commented out to get tests working for travis. This test expects an actual connection to reset
				//  to after the test has finished. In reality we don't need to reset the connection here, we just don't want to
				// create interacting tests.
				//
				// existingRedis := CacheConn()
				// // defer statments are LIFO, so this is run after the mock connection is closed.
				// defer SetConnection(existingRedis.Conn)

				mock := redigomock.NewConn()
				mock.Command("PING")

				c := SetConnection(mock)
				_, err := c.Conn().Do("PING")

				mock.Close()
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})

type Pool interface {
	Get() redis.Conn
}
