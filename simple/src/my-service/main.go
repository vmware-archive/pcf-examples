package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	// todo: parse from yaml
	port := "8000"

	http.HandleFunc("/", handler)

	println(fmt.Sprintf("Listing on %v", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
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
