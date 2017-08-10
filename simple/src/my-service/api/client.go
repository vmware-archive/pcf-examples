package api

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"my-service/db"
	"net/http"
)

type ClientAPI interface {
	GetKeyHandler(response http.ResponseWriter, request *http.Request, ps httprouter.Params)
	PutKeyHandler(response http.ResponseWriter, request *http.Request, ps httprouter.Params)
}

type clientAPI struct {
	store db.KVStore
}

func NewClientAPI(store db.KVStore) ClientAPI {
	return &clientAPI{
		store: store,
	}
}

func (client *clientAPI) GetKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: auth check
	key := params.ByName("key")
	bucketName := params.ByName("bucket_name")
	//todo: handle Get error
	value, _ := client.store.Get(bucketName, key)
	if value == nil {
		response.WriteHeader(404)
	} else {
		response.Write(value)
	}
}

func (client *clientAPI) PutKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: auth check
	key := params.ByName("key")
	bucketName := params.ByName("bucket_name")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(400)
	} else {
		//todo: handle Put error
		client.store.Put(bucketName, key, data)
	}
}
