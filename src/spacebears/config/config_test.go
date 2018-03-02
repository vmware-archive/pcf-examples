package config_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "spacebears/config"
)

var _ = Describe("Config", func() {
	Context("config parsing", func() {
		BeforeEach(func() {
			os.Clearenv()
		})

		It("parses config from environment", func() {
			os.Setenv("ADMIN_USERNAME", "bob")
			os.Setenv("ADMIN_PASSWORD", "abc123")
			os.Setenv("PORT", "9001")
			os.Setenv("DB_FILE", "custom.db")
			os.Setenv("BUCKETS", "{\"some\": \"json\"}")

			c, err := Parse()
			Expect(err).To(BeNil())
			Expect(c.AdminUsername).To(Equal("bob"))
			Expect(c.AdminPassword).To(Equal("abc123"))
			Expect(c.Port).To(Equal(9001))
			Expect(c.DBFile).To(Equal("custom.db"))
			Expect(c.Buckets).To(Equal("{\"some\": \"json\"}"))
		})

		It("check for required password", func() {
			_, err := Parse()
			Expect(err).NotTo(BeNil())
		})
	})
})
