package testutil_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/HomesNZ/go-common/testutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MockServer", func() {
	It("counts the number of requests handled", func() {
		s := NewMockFixtureServer()

		for i := int64(1); i <= 3; i++ {
			_, err := http.Get(s.URL)
			Expect(err).NotTo(HaveOccurred())
			Expect(s.RequestCount).To(Equal(i))
		}
	})

	It("resets the number of requests handled", func() {
		s := NewMockFixtureServer()

		for i := int64(1); i <= 3; i++ {
			_, err := http.Get(s.URL)
			Expect(err).NotTo(HaveOccurred())
			Expect(s.RequestCount).To(Equal(i))
		}

		s.ResetRequestCount()
		Expect(s.RequestCount).To(Equal(int64(0)))
	})

	It("responds with the fixture and status", func() {
		fixture := []byte("test fixture")
		status := http.StatusOK

		s := NewMockFixtureServer()
		s.Fixture = fixture
		s.Status = status

		resp, err := http.Get(s.URL)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(status))

		b, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(Equal(fixture))

		By("changing the fixture and status")
		fixture = []byte("Nothing to see here...")
		status = http.StatusNotFound
		s.Fixture = fixture
		s.Status = status

		resp, err = http.Get(s.URL)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(status))

		b, err = ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(Equal(fixture))
	})
})
