package api

import (
	"github.com/julienschmidt/httprouter"
	"my-service/db"
	"net/http"
)

type AdminAPI struct {
	store db.KVStore
}

func NewAdminAPI(store db.KVStore) *AdminAPI {
	return &AdminAPI{
		store: store,
	}
}

func (admin *AdminAPI) CreateBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: auth check
	//todo: input validation. bucket name empyty, etc

	//create credentials...?

	bucketName := params.ByName("bucket_name")
	if bucketName == "" {
		response.WriteHeader(400)
		return
	}
	err := admin.store.CreateBucket(bucketName)
	if err != nil {
		response.WriteHeader(500)
	}
}

func (admin *AdminAPI) DeleteBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: auth check
	bucketName := params.ByName("bucket_name")
	if bucketName == "" {
		response.WriteHeader(400)
		return
	}
	err := admin.store.DeleteBucket(bucketName)
	if err != nil {
		response.WriteHeader(500)
	}
}
