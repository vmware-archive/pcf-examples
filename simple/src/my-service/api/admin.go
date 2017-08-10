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

func (client *AdminAPI) CreateBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: auth check
}

func (client *AdminAPI) DeleteBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: auth check
}
