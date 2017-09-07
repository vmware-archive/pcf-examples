package api_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "my-service/api"
	"my-service/db/dbfakes"
)

var _ = Describe("Admin", func() {
	var mydb *dbfakes.FakeKVStore
	var adminAdpi *AdminAPI
	var logs *bytes.Buffer

	BeforeEach(func() {
		mydb = new(dbfakes.FakeKVStore)
		logs = &bytes.Buffer{}
		logger := log.New(logs, "", log.LstdFlags)

		adminAdpi = NewAdminAPI("my-admin", "my-admin-pass", mydb, logger)
	})

	Context("CreateBucketHandler", func() {
		It("success", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}

			adminAdpi.CreateBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(200))
			Expect(mydb.CreateBucketCallCount()).To(Equal(1))
			Expect(mydb.CreateBucketArgsForCall(0)).To(Equal("my_new_bucket"))
		})

		It("adds new bucket metadata", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}

			adminAdpi.CreateBucketHandler(myResponse, myRequest, myParams)

			Expect(mydb.PutCallCount()).To(Equal(1))

			putBucket, putKey, putValue := mydb.PutArgsForCall(0)
			Expect(putBucket).To(Equal("metadata"))
			Expect(putKey).To(Equal("my_new_bucket"))

			parsedPutValue := map[string]interface{}{}
			err := json.Unmarshal(putValue, &parsedPutValue)

			Expect(err).To(BeNil())
			Expect(parsedPutValue["credentials"]).To(HaveLen(0))
		})

		It("failure return 500", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.CreateBucketReturns(errors.New("Failed"))

			adminAdpi.CreateBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("validates bucket_name", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
			}

			adminAdpi.CreateBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.CreateBucketCallCount()).To(Equal(0))
		})

		It("returns error if bucket already exists", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.BucketExistsReturns(true)

			adminAdpi.CreateBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
			Expect(mydb.CreateBucketCallCount()).To(Equal(0))
		})
	})

	Context("CreateBucketCredsHandler", func() {
		It("adds credentials to metadata", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob", "password":"monkey123"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{"credentials":[{"username":"existing","password":"existing_pass"}]}`), nil)

			adminAdpi.CreateBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(mydb.PutCallCount()).To(Equal(1))

			putBucket, putKey, putValue := mydb.PutArgsForCall(0)
			Expect(putBucket).To(Equal("metadata"))
			Expect(putKey).To(Equal("my_new_bucket"))

			parsedPutValue := map[string]interface{}{}
			err := json.Unmarshal(putValue, &parsedPutValue)

			Expect(err).To(BeNil())
			Expect(parsedPutValue["credentials"]).To(HaveLen(2))
			Expect(parsedPutValue["credentials"]).To(Equal(
				[]interface{}{
					map[string]interface{}{"username": "existing", "password": "existing_pass"},
					map[string]interface{}{"username": "bob", "password": "monkey123"},
				},
			))

			Expect(myResponse.Code).To(Equal(200))
		})

		It("changes password putting existing creds", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "existing", "password":"new_existing_pass"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{"credentials":[{"username":"existing","password":"existing_pass"}]}`), nil)

			adminAdpi.CreateBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(mydb.PutCallCount()).To(Equal(1))

			putBucket, putKey, putValue := mydb.PutArgsForCall(0)
			Expect(putBucket).To(Equal("metadata"))
			Expect(putKey).To(Equal("my_new_bucket"))

			parsedPutValue := map[string]interface{}{}
			err := json.Unmarshal(putValue, &parsedPutValue)

			Expect(err).To(BeNil())
			Expect(parsedPutValue["credentials"]).To(HaveLen(1))
			Expect(parsedPutValue["credentials"]).To(Equal(
				[]interface{}{
					map[string]interface{}{"username": "existing", "password": "new_existing_pass"},
				},
			))

			Expect(myResponse.Code).To(Equal(200))
		})

		It("returns error on bucket failure", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob", "password":"monkey123"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns(nil, errors.New("db failed"))

			adminAdpi.CreateBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("returns error on failure to marshalling from db", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob", "password":"monkey123"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{`), nil)

			adminAdpi.CreateBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("returns error on storage failure", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob", "password":"monkey123"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{"credentials":[{"username":"existing","password":"existing_pass"}]}`), nil)
			mydb.PutReturns(errors.New("db storage failure"))

			adminAdpi.CreateBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("returns error on bad incoming json body", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{"credentials":[{"username":"existing","password":"existing_pass"}]}`), nil)

			adminAdpi.CreateBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
		})
	})

	Context("DeleteBucketCredsHandler", func() {
		It("removes credentials from metadata - multiple existing", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(
				`{"credentials":[
					{"username":"existing","password":"existing_pass"},
					{"username": "bob", "password":"monkey123"},
					{"username":"other_existing","password":"other_existing_pass"}
				]}`,
			), nil)

			adminAdpi.DeleteBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(mydb.PutCallCount()).To(Equal(1))

			putBucket, putKey, putValue := mydb.PutArgsForCall(0)
			Expect(putBucket).To(Equal("metadata"))
			Expect(putKey).To(Equal("my_new_bucket"))

			parsedPutValue := map[string]interface{}{}
			err := json.Unmarshal(putValue, &parsedPutValue)

			Expect(err).To(BeNil())
			Expect(parsedPutValue["credentials"]).To(HaveLen(2))
			Expect(parsedPutValue["credentials"]).To(Equal(
				[]interface{}{
					map[string]interface{}{"username": "existing", "password": "existing_pass"},
					map[string]interface{}{"username": "other_existing", "password": "other_existing_pass"},
				},
			))

			Expect(myResponse.Code).To(Equal(200))
		})

		It("removes credentials from metadata - single existing", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{"credentials":[{"username": "bob", "password":"monkey123"}]}`), nil)

			adminAdpi.DeleteBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(mydb.PutCallCount()).To(Equal(1))

			putBucket, putKey, putValue := mydb.PutArgsForCall(0)
			Expect(putBucket).To(Equal("metadata"))
			Expect(putKey).To(Equal("my_new_bucket"))

			parsedPutValue := map[string]interface{}{}
			err := json.Unmarshal(putValue, &parsedPutValue)

			Expect(err).To(BeNil())
			Expect(parsedPutValue["credentials"]).To(HaveLen(0))

			Expect(myResponse.Code).To(Equal(200))
		})

		It("returns error on bucket failure", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob", "password":"monkey123"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns(nil, errors.New("db failed"))

			adminAdpi.DeleteBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("returns error on failure to marshalling from db", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "bob", "password":"monkey123"}`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{`), nil)

			adminAdpi.DeleteBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("returns error on bad incoming json body", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Body = ioutil.NopCloser(bytes.NewBufferString(`{`))
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.GetReturns([]byte(`{"credentials":[{"username":"existing","password":"existing_pass"}]}`), nil)

			adminAdpi.DeleteBucketCredsHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(400))
		})
	})

	Context("DeleteBucketHandler", func() {
		It("success", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}

			adminAdpi.DeleteBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(200))
			Expect(mydb.DeleteBucketCallCount()).To(Equal(1))
			Expect(mydb.DeleteBucketArgsForCall(0)).To(Equal("my_new_bucket"))
		})

		It("failure return 500", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.DeleteBucketReturns(errors.New("Failed"))

			adminAdpi.DeleteBucketHandler(myResponse, myRequest, myParams)

			Expect(myResponse.Code).To(Equal(500))
		})

		It("validates bucket_name", func() {
			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
			}

			adminAdpi.DeleteBucketHandler(myResponse, myRequest, myParams)

			Expect(mydb.DeleteBucketCallCount()).To(Equal(0))
			Expect(myResponse.Code).To(Equal(400))
		})
	})

	Context("AdminAuthFilter", func() {
		It("returns 401 without authentication header", func() {
			nextRequestInvoked := false
			nextRequest := func(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
				nextRequestInvoked = true
			}
			filtered := adminAdpi.AdminAuthFilter(nextRequest)

			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{}

			filtered(myResponse, myRequest, myParams)

			Expect(nextRequestInvoked).To(BeFalse())
			Expect(myResponse.Code).To(Equal(401))
		})

		It("returns 401 with improper auth header", func() {
			nextRequestInvoked := false
			nextRequest := func(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
				nextRequestInvoked = true
			}
			filtered := adminAdpi.AdminAuthFilter(nextRequest)

			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Header.Set(
				"authorization",
				fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("foo:bar"))),
			)

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{}

			filtered(myResponse, myRequest, myParams)

			Expect(nextRequestInvoked).To(BeFalse())
			Expect(myResponse.Code).To(Equal(401))
		})

		It("lets through properly auth'd requests", func() {
			nextRequestInvoked := false
			nextRequest := func(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
				nextRequestInvoked = true
				response.WriteHeader(200)
			}
			filtered := adminAdpi.AdminAuthFilter(nextRequest)

			myRequest := httptest.NewRequest("GET", "https://example.com", nil)
			myRequest.Header.Set(
				"Authorization",
				fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("my-admin:my-admin-pass"))),
			)

			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{}

			filtered(myResponse, myRequest, myParams)

			Expect(nextRequestInvoked).To(BeTrue())
			Expect(myResponse.Code).To(Equal(200))
		})
	})
})
