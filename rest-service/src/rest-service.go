package main

import (
	"fmt"
	"net/http"
)

func helloHandler(response http.ResponseWriter, req *http.Request) {
	msg := "Hello World"
	fmt.Fprintln(response, msg)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", helloHandler)
	http.ListenAndServe(":9090", router)
}
