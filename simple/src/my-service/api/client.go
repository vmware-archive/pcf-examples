package api

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"my-service/db"
	"net/http"
)

type ClientAPI struct {
	store db.KVStore
}

func NewClientAPI(store db.KVStore) *ClientAPI {
	return &ClientAPI{
		store: store,
	}
}

func (client *ClientAPI) GetKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

func (client *ClientAPI) PutKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
