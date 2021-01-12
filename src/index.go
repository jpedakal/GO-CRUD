package main

import (
	handler "api"
	"net/http"
	database "utils"
)
func main() {
  
	database.Connect()
	http.HandleFunc("/insert", handler.PostData)
	http.HandleFunc("/fetch", handler.GetData)
	http.ListenAndServe(":8000", nil)
}
