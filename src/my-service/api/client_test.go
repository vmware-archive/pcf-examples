package api_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"encoding/base64"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "my-service/api"
	"my-service/db"
	"my-service/db/dbfakes"
)

var _ = Describe("Client", func() {
	var mydb *dbfakes.FakeKVStore
	var logs *bytes.Buffer
	var logger *log.Logger

	BeforeEach(func() {
		mydb = new(dbfakes.FakeKVStore)
		logs = &bytes.Buffer{}
		logger = log.New(logs, "", log.LstdFlags)
	})

	Context("PutKeyHandler", func() {
		It("success", func() {
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			data := ioutil.NopCloser(strings.NewReader("my new value"))
			myRequest := httptest.NewRequest("", "/", data)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.PutKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(200))

			Expect(mydb.PutCallCount()).To(Equal(1))
			givenBucketName, givenKey, givenValue := mydb.PutArgsForCall(0)
			Expect(givenBucketName).To(Equal("myfirstbucket"))
			Expect(givenKey).To(Equal("mykey"))
			Expect(givenValue).To(Equal([]byte("my new value")))
		})

		It("401 when no auth header", func() {
			data := ioutil.NopCloser(strings.NewReader("my new value"))
			myRequest := &http.Request{
				Body: data,
			}
			myResponse := httptest.NewRecorder()

			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.PutKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(401))
		})

		It("401 when bad auth", func() {
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			data := ioutil.NopCloser(strings.NewReader("my new value"))
			myRequest := httptest.NewRequest("", "/", data)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:badpass")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.PutKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(401))
		})

		It("input validation failed", func() {
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
				{Key: "key", Value: ""},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.PutKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.PutCallCount()).To(Equal(0))
		})

		It("body is garbage", func() {
			data := ioutil.NopCloser(&ErrorReader{})
			myRequest := httptest.NewRequest("", "/", data)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			myResponse := httptest.NewRecorder()
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)

			client := NewClientAPI(mydb, logger)
			client.PutKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.PutCallCount()).To(Equal(0))
		})

		It("put throws error", func() {
			mydb.PutReturns(errors.New("something bad happen disk is corrupted"))
			data := ioutil.NopCloser(strings.NewReader("my new value"))
			myRequest := httptest.NewRequest("", "/", data)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			myResponse := httptest.NewRecorder()
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)

			client := NewClientAPI(mydb, logger)
			client.PutKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(500))
		})
	})

	Context("ListBucketHandler", func() {
		It("success", func() {
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			mydb.ListReturns([]db.KeyValue{
				{Key: []byte("foo"), Value: []byte("bar")},
				{Key: []byte("baz"), Value: []byte("qux")},
			}, nil)

			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.ListBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(200))

			Expect(mydb.ListCallCount()).To(Equal(1))
			Expect(myResponse.Body.String()).To(Equal(`{"baz":"qux","foo":"bar"}`))
		})

		It("returns 500 on db failure", func() {
			mydb.GetReturns([]byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			mydb.ListReturns(nil, errors.New("db failure"))

			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.ListBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
			Expect(mydb.ListCallCount()).To(Equal(1))
		})

		It("bad credentials returns auth failure", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"hjljhkhkhj"}]}`), nil)
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.ListBucketHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(401))
		})

		It("400 when no bucketName", func() {
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.ListBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
		})
	})

	Context("GetKeyHandler", func() {
		It("success", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			mydb.GetReturnsOnCall(1, []byte("myvalue"), nil)

			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.GetKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(200))
			Expect(myResponse.Body.String()).To(Equal("myvalue"))
			Expect(myResponse.Header()["Content-Type"][0]).To(HavePrefix("text/plain"))

			Expect(mydb.GetCallCount()).To(Equal(2))
			givenBucketName, givenKey := mydb.GetArgsForCall(1)
			Expect(givenBucketName).To(Equal("myfirstbucket"))
			Expect(givenKey).To(Equal("mykey"))
		})

		It("input validation failed", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
				{Key: "key", Value: ""},
			}

			client := NewClientAPI(mydb, logger)
			client.GetKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.GetCallCount()).To(Equal(0))
		})

		It("key not found", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			mydb.GetReturnsOnCall(1, nil, nil)
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.GetKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(404))
		})

		It("get throws error", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			mydb.GetReturnsOnCall(1, nil, errors.New("something bad happen disk is corrupted"))
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.GetKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(500))
		})

		It("bad credentials returns auth failure", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"hjljhkhkhj"}]}`), nil)
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.GetKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(401))
		})
	})

	Context("DeleteKeyHandler", func() {
		It("success", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)

			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			myResponse := httptest.NewRecorder()

			client := NewClientAPI(mydb, logger)
			client.DeleteKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(200))

			Expect(mydb.DeleteCallCount()).To(Equal(1))
			givenBucketName, givenKey := mydb.DeleteArgsForCall(0)
			Expect(givenBucketName).To(Equal("myfirstbucket"))
			Expect(givenKey).To(Equal("mykey"))
		})

		It("input validation failed", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
				{Key: "key", Value: ""},
			}

			client := NewClientAPI(mydb, logger)
			client.DeleteKeyHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.GetCallCount()).To(Equal(0))
		})

		It("db throws error", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"abc123"}]}`), nil)
			mydb.DeleteReturnsOnCall(0, errors.New("something bad happen disk is corrupted"))
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.DeleteKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(500))
		})

		It("bad credentials returns auth failure", func() {
			mydb.GetReturnsOnCall(0, []byte(`{"credentials":[{"username":"user","password":"hjljhkhkhj"}]}`), nil)
			myRequest := httptest.NewRequest("", "/", nil)
			myRequest.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:abc123")))

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewClientAPI(mydb, logger)
			client.DeleteKeyHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(401))
		})
	})
})

type ErrorReader struct{}

func (r *ErrorReader) Read(b []byte) (n int, err error) {
	return 0, errors.New("This reader always fails")
}
