package api_test

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"

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

			//todo: check bucket created
		})
	})
})
