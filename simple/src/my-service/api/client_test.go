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
	"my-service/db/dbfakes"
)

var _ = Describe("Client", func() {
	var mydb *dbfakes.FakeKVStore

	BeforeEach(func() {
		mydb = new(dbfakes.FakeKVStore)
	})

	Context("GetKeyHandler", func() {
		It("success", func() {
			mydb := &dbfakes.FakeKVStore{}
			mydb.GetReturns([]byte("myvalue"), nil)
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

			Expect(mydb.GetCallCount()).To(Equal(1))
			givenBucketName, givenKey := mydb.GetArgsForCall(0)
			Expect(givenBucketName).To(Equal("myfirstbucket"))
			Expect(givenKey).To(Equal("mykey"))
		})

		It("input validation failed", func() {
			mydb := &dbfakes.FakeKVStore{}

			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
				{Key: "key", Value: ""},
			}

			client := NewClientAPI(mydb)
			client.GetKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.GetCallCount()).To(Equal(0))
		})

		It("key not found", func() {
			mydb.GetReturns(nil, nil)
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

		It("get throws error", func() {
			mydb.GetReturns(nil, errors.New("something bad happen disk is corrupted"))
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb)
			client.GetKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(500))
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

			Expect(mydb.PutCallCount()).To(Equal(1))
			givenBucketName, givenKey, givenValue := mydb.PutArgsForCall(0)
			Expect(givenBucketName).To(Equal("myfirstbucket"))
			Expect(givenKey).To(Equal("mykey"))
			Expect(givenValue).To(Equal([]byte("my new value")))
		})

		It("input validation failed", func() {
			mydb := &dbfakes.FakeKVStore{}
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
				{Key: "key", Value: ""},
			}

			client := NewClientAPI(mydb)
			client.PutKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.PutCallCount()).To(Equal(0))

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
			Expect(mydb.PutCallCount()).To(Equal(0))
		})

		It("put throws error", func() {
			mydb.PutReturns(errors.New("something bad happen disk is corrupted"))
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
			Expect(myResponse.Code).To(Equal(500))
		})
	})
})

type ErrorReader struct{}

func (r *ErrorReader) Read(b []byte) (n int, err error) {
	return 0, errors.New("This reader always fails")
}
