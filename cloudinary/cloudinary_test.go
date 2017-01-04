package cloudinary_test

import (
	"net/http"
	"os"
	"strings"

	. "github.com/HomesNZ/go-common/cloudinary"

	"github.com/HomesNZ/go-common/testutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cloudinary", func() {
	os.Setenv("CLOUDINARY_API_KEY", "key")
	os.Setenv("CLOUDINARY_API_SECRET", "secret")
	os.Setenv("CLOUDINARY_CLOUD_NAME", "testcloud")

	Describe(".UploadImage", func() {
		It("performs a POST request to the Cloudinary API", func() {
			resp := []byte(`{"public_id":"tests/test_file","version":1369431906,"format":"png","resource_type":"image"}`)

			server := testutil.NewMockFixtureServer()
			server.Status = http.StatusOK
			server.Fixture = resp
			defer server.Close()

			s := Service()
			err := s.UploadURI(server.URL)
			Expect(err).NotTo(HaveOccurred())

			u, err := s.UploadImage("test", strings.NewReader(""))
			Expect(err).NotTo(HaveOccurred())

			expectedURL := "http://res.cloudinary.com/testcloud/image/upload/test_"

			Expect(server.RequestCount).To(Equal(int64(1)))
			// The url ends with a unix nano timestamp, so we can only deterministically match the prefix.
			Expect(u).To(HavePrefix(expectedURL))
		})
	})
})
