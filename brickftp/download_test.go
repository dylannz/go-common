package brickftp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/HomesNZ/go-common/brickftp"
	"path"
)

var _ = Describe("Download", func() {
	Context("zip archive", func() {
		It("extracts a zip file", func() {
			val, err := brickftp.UnpackZIP("test/Archive.zip")
			Expect(err).NotTo(HaveOccurred())

			Expect(path.Base(val)).To(Equal("Archive"))
		})
	})
})