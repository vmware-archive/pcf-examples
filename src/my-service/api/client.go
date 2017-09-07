package api

import (
	"io/ioutil"
	"log"
	"my-service/db"
	"net/http"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"my-service/models"
)

type ClientAPI struct {
	store  db.KVStore
	logger *log.Logger
}

func NewClientAPI(store db.KVStore, logger *log.Logger) *ClientAPI {
	return &ClientAPI{
		store:  store,
		logger: logger,
	}
}

func (client *ClientAPI) PutKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := params.ByName("key")
	bucketName := params.ByName("bucket_name")
	if bucketName == "" || key == "" {
		response.WriteHeader(400)
		return
	}

	if !client.checkAuth(response, request, bucketName) {
		return
	}

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		client.logger.Print(err)
		response.WriteHeader(400)
		return
	} else {
		err := client.store.Put(bucketName, key, data)
		if err != nil {
			client.logger.Print(err)
			response.WriteHeader(500)
			return
		}
	}
}

func (client *ClientAPI) GetKeyHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := params.ByName("key")
	bucketName := params.ByName("bucket_name")
	if bucketName == "" || key == "" {
		response.WriteHeader(400)
		return
	}

	if !client.checkAuth(response, request, bucketName) {
		return
	}

	value, err := client.store.Get(bucketName, key)
	if err != nil {
		client.logger.Print(err)
		response.WriteHeader(500)
		return
	} else if value == nil {
		response.WriteHeader(404)
		return
	} else {
		response.Write(value)
	}
}

func (client *ClientAPI) checkAuth(response http.ResponseWriter, request *http.Request, bucketName string) bool {
	response.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	username, password, authOK := request.BasicAuth()
	if authOK == false {
		http.Error(response, "Not authorized", 401)
		return false
	}

	rawMetadata, err := client.store.Get("metadata", bucketName)
	if err != nil {
		client.logger.Print(err)
		response.WriteHeader(500)
		return false
	}
	metadata := models.BucketMetadata{}
	err = json.Unmarshal(rawMetadata, &metadata)
	if err != nil {
		client.logger.Print(err)
		response.WriteHeader(500)
		return false
	}

	for _, creds := range metadata.Credentials {
		if creds.Username == username && creds.Password == password {
			return true
		}
	}

	http.Error(response, "Not authorized", 401)
	return false
}
