package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	go heartbeat()
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)

	println(fmt.Sprintf("Listing on %v", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func heartbeat() {
	time.Sleep(1 * time.Second)
	println(fmt.Sprintf("New heartbeat %v", time.Now()))
	heartbeat()
}

func handler(response http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		println("Error reading request", err)
		return
	}

	requestBody := string(bodyBytes)
	responseBody := "    No incoming request body!"
	if requestBody != "" {
		responseBody = requestBody
	} else {
		response.WriteHeader(422)
	}
	response.Write([]byte(fmt.Sprintf("Echo server given:\n%v\n", responseBody)))
}
