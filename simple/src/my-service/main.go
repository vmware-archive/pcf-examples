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

	//todo: make path configurable
	mydb, err := db.NewDB("my.db")
	client := api.NewClientAPI(mydb)
	if err != nil {
		println(fmt.Sprintf("%v", err))
		syscall.Exit(1)
	}

	router := httprouter.New()
	router.GET("/api/:bucket_name/:key", client.GetKeyHandler)
	router.PUT("/api/:bucket_name/:key", client.PutKeyHandler)
	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Port), router)

	if err != nil {
		println(fmt.Sprintf("%v", err))
		syscall.Exit(1)
	}
}
