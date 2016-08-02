package brickftp_test

import (
	. "github.com/HomesNZ/go-common/brickftp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Brickftp", func() {
	DescribeTable("#Permitted",
		func(action string, expected bool) {
			webhook := BrickFTP{
				Action: action,
			}
			Expect(webhook.Permitted()).To(Equal(expected))
		},
		Entry("is true with the 'create' action ", "create", true),
		Entry("is false with the 'update' action", "update", false),
		Entry("is false with an empty action", "", false),
	)
})
