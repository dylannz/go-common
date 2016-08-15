package zip_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path"
	"github.com/HomesNZ/go-common/zip"
	"os"
)

var _ = Describe("Download", func() {
	Context("zip archive", func() {
		It("extracts a zip file", func() {
			val, err := zip.Unpack("test/Archive.zip")
			Expect(err).NotTo(HaveOccurred())

			Expect(path.Base(val)).To(Equal("Archive"))

			err = os.RemoveAll(val)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})