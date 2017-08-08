package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/julienschmidt/httprouter"
	. "my-service/api"
	"my-service/db"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Client", func() {

	BeforeEach(func() {

	})

	It("tests GetKeyHandler success", func() {

		mydb := db.NewDB()
		mydb.Put("myfirstbucket", "mykey", "myvalue")

		myRequest := &http.Request{}
		myResponse := httptest.NewRecorder()
		myParams := httprouter.Params{
			{Key: "bucket_name", Value: "myfirstbucket"},
			{Key: "key", Value: "mykey"},
		}

		client := NewClientAPI(mydb)
		client.GetKeyHandler(myResponse, myRequest, myParams)
		Expect(myResponse.Body.String()).To(Equal("myvalue"))
	})

})
