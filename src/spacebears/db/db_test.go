package db_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"

	. "spacebears/db"
)

var _ = Describe("Db", func() {
	var db KVStore
	var tempDir string

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "")
		Expect(err).To(BeNil())

		dbPath := fmt.Sprintf("%s%c%s", tempDir, os.PathSeparator, "test.db")
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

	It("list bucket contents", func() {
		err := db.Put("mybucket", "mykey0", []byte("myvalue0"))
		Expect(err).To(BeNil())

		err = db.Put("mybucket", "mykey1", []byte("myvalue1"))
		Expect(err).To(BeNil())

		err = db.Put("mybucket", "mykey2", []byte("myvalue2"))
		Expect(err).To(BeNil())

		list, err := db.List("mybucket")
		Expect(err).To(BeNil())
		Expect(list).To(HaveLen(3))
		expectedList := []KeyValue{
			{Key: []byte("mykey0"), Value: []byte("myvalue0")},
			{Key: []byte("mykey1"), Value: []byte("myvalue1")},
			{Key: []byte("mykey2"), Value: []byte("myvalue2")},
		}
		for i, kv := range list {
			Expect(expectedList[i]).To(Equal(kv))
		}
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

	It("creats funky named bucket", func() {
		err := db.CreateBucket("∴_///3asdfa33332:::")

		Expect(err).To(BeNil())

		err = db.CreateBucket(" ")

		Expect(err).To(BeNil())
	})

	It("create bucket is idempotent", func() {
		err := db.CreateBucket("foo")
		Expect(err).To(BeNil())

		err = db.CreateBucket("foo")
		Expect(err).To(BeNil())
	})

	It("delete removes key", func() {
		err := db.Put("mybucket", "mykey", []byte("myvalue"))
		Expect(err).To(BeNil())

		returnedValue, err := db.Get("mybucket", "mykey")

		Expect(err).To(BeNil())
		Expect(returnedValue).To(Equal([]byte("myvalue")))

		err = db.Delete("mybucket", "mykey")
		Expect(err).To(BeNil())

		returnedValue, err = db.Get("mybucket", "mykey")
		Expect(err).To(BeNil())
		Expect(returnedValue).To(BeNil())
	})

	It("bucket exists", func() {
		exists := db.BucketExists("foo")
		Expect(exists).To(BeFalse())

		err := db.CreateBucket("foo")
		Expect(err).To(BeNil())

		exists = db.BucketExists("foo")
		Expect(exists).To(BeTrue())
	})
})
