package util_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "my-service/util"
)

var _ = Describe("CredsGenerator", func() {

	var mycredsGenerator CredsGenerator

	BeforeEach(func() {
		mycredsGenerator = NewCredsGenerator()
	})

	It("success", func() {
		userName, password, err := mycredsGenerator.Generate()

		Expect(userName).ToNot(Equal(""))
		Expect(password).ToNot(Equal(""))
		Expect(err).To(BeNil())
		Expect(userName).ToNot(Equal(password))

	})

})
