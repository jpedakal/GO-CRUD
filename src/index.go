package main

import (
	handler "api"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handler.HelloHandler)
	http.ListenAndServe(":8000", nil)
}
