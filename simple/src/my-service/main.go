package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"

	"my-service/api"
	"my-service/config"
	"my-service/db"
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

	client := api.NewClientAPI(mydb)
	admin := api.NewAdminAPI(c.AdminUsername, c.AdminPassword, mydb, logger)

	router := httprouter.New()
	router.POST("/api/bucket/:bucket_name", admin.AdminAuthFilter(admin.CreateBucketHandler))
	router.PUT("/api/bucket/:bucket_name/credentials", admin.AdminAuthFilter(admin.CreateBucketCredsHandler))
	router.DELETE("/api/bucket/:bucket_name/credentials", admin.DeleteBucketCredsHandler)
	router.DELETE("/api/bucket/:bucket_name", admin.AdminAuthFilter(admin.DeleteBucketHandler))
	router.GET("/api/bucket/:bucket_name/:key", client.GetKeyHandler)
	router.PUT("/api/bucket/:bucket_name/:key", client.PutKeyHandler)

	logger.Println(fmt.Sprintf("Listning on port %v", c.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Port), router)
	if err != nil {
		logger.Fatal(err)
	}
}
