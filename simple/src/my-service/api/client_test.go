package api_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "my-service/api"
	"my-service/db"
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
