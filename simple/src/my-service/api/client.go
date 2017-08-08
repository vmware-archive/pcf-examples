package api

import (
	"my-service/db"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"io/ioutil"
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
	key := params.ByName("key")
	bucketName:= params.ByName("bucket_name")
	println(fmt.Sprintf("bucket: %v", bucketName))
	println(fmt.Sprintf("key: %v", key))
	value := client.store.Get(bucketName, key)
	println(fmt.Sprintf("value: %v", value))
	// todo: json, probably
	serialized := fmt.Sprintf("%v", value)
	response.Write([]byte(string(serialized)))
}

func (client *clientAPI) PutKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := params.ByName("key")
	bucketName:= params.ByName("bucket_name")
	println(fmt.Sprintf("bucket: %v", bucketName))
	println(fmt.Sprintf("key: %v", key))
	data, _ := ioutil.ReadAll(request.Body)
	client.store.Put( bucketName, key, string(data))

}