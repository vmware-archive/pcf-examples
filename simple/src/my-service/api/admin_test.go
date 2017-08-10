package api_test

import (
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

	BeforeEach(func() {
		mydb = new(dbfakes.FakeKVStore)
	})

	Context("CreateBucketHandler", func() {
		It("success", func() {
			myRequest := &http.Request{}
			myResponse := httptest.NewRecorder()
			myParams := httprouter.Params{
				{Key: "bucket_name", Value: "myfirstbucket"},
				{Key: "key", Value: "mykey"},
			}

			client := NewAdminAPI(mydb)
			client.CreateBucketHandler(myResponse, myRequest, myParams)
			Expect(myResponse.Code).To(Equal(200))

			Expect(mydb.CreateBucketCallCount()).To(Equal(1))
			Expect(mydb.CreateBucketArgsForCall(0)).To(Equal("my_new_bucket"))
		})
	})
})
