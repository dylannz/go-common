package util_test

import (
	. "github.com/HomesNZ/data-import/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	DescribeTable("#Contains",
		func(key string, expected bool) {
			set := NewSet("exists", "also_exists")
			Expect(set.Contains(key)).To(Equal(expected))
		},
		Entry("returns true when the key exists", "exists", true),
		Entry("returns false when the key doesn't exist", "doesnt_exist", false),
		Entry("returns false with an empty key", "", false),
	)
})
