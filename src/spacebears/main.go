package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"

	"spacebears/api"
	"spacebears/config"
	"spacebears/db"
)

var logger *log.Logger

func main() {
	logger = log.New(os.Stderr, "", log.LstdFlags)

	c, err := config.Parse()
	if err != nil {
		logger.Fatal(err)
	}

	mydb, err := db.NewDB(c.DBFile)
	if err != nil {
		logger.Fatal(err)
	}
	err = mydb.CreateBucket("metadata")
	if err != nil {
		logger.Fatal(err)
	}

	client := api.NewClientAPI(mydb, logger)
	admin := api.NewAdminAPI(c.AdminUsername, c.AdminPassword, mydb, logger)

	router := httprouter.New()
	router.GET("/health", func(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		response.Write([]byte("healthy"))
	})

	router.POST("/api/admin/bucket/:bucket_name", admin.AdminAuthFilter(admin.CreateBucketHandler))
	router.PUT("/api/admin/bucket/:bucket_name/credentials", admin.AdminAuthFilter(admin.CreateBucketCredsHandler))
	router.DELETE("/api/admin/bucket/:bucket_name/credentials", admin.DeleteBucketCredsHandler)
	router.DELETE("/api/admin/bucket/:bucket_name", admin.AdminAuthFilter(admin.DeleteBucketHandler))

	router.PUT("/api/bucket/:bucket_name/:key", client.PutKeyHandler)
	router.GET("/api/bucket/:bucket_name/", client.ListBucketHandler)
	router.GET("/api/bucket/:bucket_name/:key", client.GetKeyHandler)
	router.DELETE("/api/bucket/:bucket_name/:key", client.DeleteKeyHandler)

	logger.Println(fmt.Sprintf("Listning on port %v", c.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Port), router)
	if err != nil {
		logger.Fatal(err)
	}
}
