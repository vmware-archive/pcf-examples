package db_test

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "my-service/db"
)

var _ = Describe("Db", func() {
	var db KVStore
	var tempDir string

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "")
		Expect(err).To(BeNil())

		dbPath := fmt.Sprintf("%s%s%s", tempDir, os.PathSeparator, "test.db")
		db, err = NewDB(dbPath)
		Expect(err).To(BeNil())

		err = db.CreateBucket("mybucket")
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		os.Remove(tempDir)
	})

	It("put and retrieve value", func() {
		err := db.Put("mybucket", "mykey", []byte("myvalue"))
		Expect(err).To(BeNil())

		returnedValue, err := db.Get("mybucket", "mykey")

		Expect(err).To(BeNil())
		Expect(returnedValue).To(Equal([]byte("myvalue")))
	})

	It("put elevates error", func() {
		db.Close()
		err := db.Put("mybucket", "mykey", []byte("myvalue"))

		Expect(err).NotTo(BeNil())
	})

	It("get elevates error", func() {
		db.Close()
		_, err := db.Get("mybucket", "mykey")

		Expect(err).NotTo(BeNil())
	})

	It("create bucket elevates error", func() {
		db.Close()
		err := db.CreateBucket("my_other_bucket")

		Expect(err).NotTo(BeNil())
	})
})
