package brickftp_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGeocoding(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Download Suite Suite")
}
