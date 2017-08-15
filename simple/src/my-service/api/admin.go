package api

import (
	"encoding/base64"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"my-service/db"
)

type AdminAPI struct {
	adminUser string
	adminPass string
	store     db.KVStore
}

func NewAdminAPI(adminUser string, adminPass string, store db.KVStore) *AdminAPI {
	return &AdminAPI{
		adminUser: adminUser,
		adminPass: adminPass,
		store:     store,
	}
}

func (admin *AdminAPI) CreateBucketHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//todo: create credentials...?

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
