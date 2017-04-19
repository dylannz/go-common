package extmath_test

import (
	"github.com/HomesNZ/go-common/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUtil(t *testing.T) {
	config.InitLogger()

	RegisterFailHandler(Fail)
	RunSpecs(t, "extmath Suite")
}
