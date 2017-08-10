package api_test

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"

	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "my-service/api"
	"my-service/db/dbfakes"
)

var _ = Describe("Admin", func() {
	var mydb *dbfakes.FakeKVStore

	BeforeEach(func() {
		mydb = new(dbfakes.FakeKVStore)
	})

	Context("CreateBucketHandler", func() {
		It("success", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}

			admin := NewAdminAPI(mydb)
			admin.CreateBucketHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(200))

			Expect(mydb.CreateBucketCallCount()).To(Equal(1))
			Expect(mydb.CreateBucketArgsForCall(0)).To(Equal("my_new_bucket"))
		})

		It("failure return 500", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}
			mydb.CreateBucketReturns(errors.New("Failed"))
			admin := NewAdminAPI(mydb)
			admin.CreateBucketHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(500))
		})

		It("validates bucket_name", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
			}
			admin := NewAdminAPI(mydb)
			admin.CreateBucketHandler(myResponse, myRequest, myParams)
			Expect(mydb.CreateBucketCallCount()).To(Equal(0))
			Expect(myResponse.Code).To(Equal(400))
		})
	})

	Context("DeleteBucketHandler", func() {
		It("success", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "my_new_bucket"},
			}

			admin := NewAdminAPI(mydb)
			admin.DeleteBucketHandler(myResponse, myRequest, myParams)
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
			admin := NewAdminAPI(mydb)
			admin.DeleteBucketHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(500))
		})

		It("validates bucket_name", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: ""},
			}
			admin := NewAdminAPI(mydb)
			admin.DeleteBucketHandler(myResponse, myRequest, myParams)
			Expect(mydb.DeleteBucketCallCount()).To(Equal(0))
			Expect(myResponse.Code).To(Equal(400))
		})
	})
})
