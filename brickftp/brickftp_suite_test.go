package brickftp_test

import (
	"github.com/HomesNZ/go-common/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBrickftp(t *testing.T) {
	config.InitLogger()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Brickftp Suite")
}
