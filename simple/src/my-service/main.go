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

	mydb := db.NewDB()
	client := api.NewClientAPI(mydb)

	router := httprouter.New()
	router.GET("/api/:bucket_name/:key", client.GetKeyHandler)
	router.PUT("/api/:bucket_name/:key", client.PutKeyHandler)
	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Port), router)

	if err != nil {
		println(fmt.Sprintf("%v", err))
		syscall.Exit(1)
	}
}
