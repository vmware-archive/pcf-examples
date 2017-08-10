package api_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "my-service/api"
	"my-service/db"
)

var _ = Describe("Client", func() {
	var mydb db.KVStore

	BeforeEach(func() {
		mydb = &memoryDb{
			data: map[string]map[string][]byte{},
		}
	})

	Context("GetKeyHandler", func() {
		It("success", func() {
			mydb.Put("myfirstbucket", "mykey", []byte("myvalue"))

			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb)
			client.GetKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(200))
			Expect(myResponse.Body.String()).To(Equal("myvalue"))
			Expect(myResponse.Header()["Content-Type"][0]).To(HavePrefix("text/plain"))
		})

		It("key not found", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb)
			client.GetKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(404))
		})
	})

	Context("PutKeyHandler", func() {
		It("success", func() {
			data := ioutil.NopCloser(strings.NewReader("my new value"))

			myRequest := &http.Request{
				Body: data,
			}
			myResponse := httptest.NewRecorder()

			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb)
			client.PutKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(200))
			Expect(mydb.Get("myfirstbucket", "mykey")).To(Equal([]byte("my new value")))

		})
		It("body is garbage", func() {
			data := ioutil.NopCloser(&ErrorReader{})

			myRequest := &http.Request{
				Body: data,
			}
			myResponse := httptest.NewRecorder()

			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb)
			client.PutKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(400))
		})
	})
})

type ErrorReader struct{}

func (r *ErrorReader) Read(b []byte) (n int, err error) {
	return 0, errors.New("This reader always fails")
}

type memoryDb struct {
	data map[string]map[string][]byte
}

func (db *memoryDb) Put(bucketName string, key string, value []byte) error {
	bucket, ok := db.data[bucketName]
	if !ok {
		bucket = make(map[string][]byte)
		db.data[bucketName] = bucket
	}
	bucket[key] = value
	return nil
}

func (db *memoryDb) Get(bucketName string, key string) ([]byte, error) {
	return db.data[bucketName][key], nil
}

func (db *memoryDb) CreateBucket(bucketName string) error {
	return nil
}

func (db *memoryDb) Close() error {
	return nil
}
