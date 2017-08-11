package main

import (
	"fmt"
	"net/http"
	"syscall"

	"github.com/julienschmidt/httprouter"

	"my-service/api"
	"my-service/config"
	"my-service/db"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		println(fmt.Sprintf("%v", err))
		syscall.Exit(1)
	}

	mydb, err := db.NewDB(c.DBFile)
	exitOnError(err)
	err = mydb.CreateBucket("metadata")
	exitOnError(err)

	client := api.NewClientAPI(mydb)
	admin := api.NewAdminAPI(mydb)

	router := httprouter.New()
	router.POST("/api/bucket/:bucket_name", admin.CreateBucketHandler)
	router.DELETE("/api/bucket/:bucket_name", admin.DeleteBucketHandler)
	router.GET("/api/bucket/:bucket_name/:key", client.GetKeyHandler)
	router.PUT("/api/bucket/:bucket_name/:key", client.PutKeyHandler)
	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Port), router)

	exitOnError(err)
}

func exitOnError(err error) {
	if err != nil {
		println(fmt.Sprintf("%v", err))
		syscall.Exit(1)
	}
}
