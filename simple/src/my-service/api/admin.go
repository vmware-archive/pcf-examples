package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"my-service/db"
	"my-service/models"
)

type AdminAPI struct {
	adminUser string
	adminPass string
	store     db.KVStore
	logger    *log.Logger
}

func NewAdminAPI(adminUser string, adminPass string, store db.KVStore, logger *log.Logger) *AdminAPI {
	return &AdminAPI{
		adminUser: adminUser,
		adminPass: adminPass,
		store:     store,
		logger:    logger,
	}
}

func (admin *AdminAPI) CreateBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucket_name")
	if bucketName == "" ||  bucketName == "metadata" {
		response.WriteHeader(400)
		return
	}
	exists := admin.store.BucketExists(bucketName)
	if exists {
		response.WriteHeader(400)
		response.Write([]byte("Bucket exists"))
		return
	}

	err := admin.store.CreateBucket(bucketName)
	if err != nil {
		log.Print(err)
		response.WriteHeader(500)
	}

	rawMetadata := models.BucketMetadata{
		Credentials: []models.BucketCredentials{},
	}
	metadata, err := json.MarshalIndent(rawMetadata, "", "")

	admin.store.Put("metadata", bucketName, []byte(metadata))
}

func (admin *AdminAPI) CreateBucketCredsHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucket_name")
	rawMetadata, err := admin.store.Get("metadata", bucketName)
	if err != nil {
		admin.logger.Print(err)
		response.WriteHeader(500)
		return
	}
	metadata := models.BucketMetadata{}
	err = json.Unmarshal(rawMetadata, &metadata)
	if err != nil {
		admin.logger.Print(err)
		response.WriteHeader(500)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		admin.logger.Print(err)
		response.WriteHeader(500)
		return
	}
	newCredentials := models.BucketCredentials{}
	err = json.Unmarshal(body, &newCredentials)
	if err != nil {
		admin.logger.Print(err)
		response.WriteHeader(400)
		return
	}

	existing := false
	for i, _ := range metadata.Credentials {
		existingCredentials := &metadata.Credentials[i]
		if existingCredentials.Username == newCredentials.Username {
			existingCredentials.Password = newCredentials.Password
			existing = true
		}
	}
	if !existing {
		metadata.Credentials = append(metadata.Credentials, newCredentials)
	}

	serializedMetadata, err := json.MarshalIndent(metadata, "", "")
	err = admin.store.Put("metadata", bucketName, serializedMetadata)
	if err != nil {
		admin.logger.Print(err)
		response.WriteHeader(500)
		return
	}
}

func (admin *AdminAPI) DeleteBucketCredsHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("todo: implement")
}

func (admin *AdminAPI) DeleteBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: remove credentials
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

func (admin *AdminAPI) AdminAuthFilter(handle httprouter.Handle) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		authHeader := request.Header.Get("Authorization")
		authEncoded := base64.StdEncoding.EncodeToString([]byte(admin.adminUser + ":" + admin.adminPass))
		if authHeader != "Basic "+authEncoded {
			response.Header().Set("WWW-Authenticate", `Basic realm="Admin"`)
			response.WriteHeader(401)
			response.Write([]byte("401 Unauthorized\n"))
		} else {
			handle(response, request, params)
		}
	}
}
