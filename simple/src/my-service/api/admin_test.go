package api_test

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "my-service/api"
	"my-service/db/dbfakes"
)

var _ = Describe("Admin", func() {
	var mydb *dbfakes.FakeKVStore
	var adminAdpi *AdminAPI

	BeforeEach(func() {
		mydb = new(dbfakes.FakeKVStore)
		adminAdpi = NewAdminAPI("my-admin", "my-admin-pass", mydb)
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

			Expect(mydb.CreateBucketCallCount()).To(Equal(0))
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
