package main

import (
	handler "api"
	"net/http"
	database "utils"
)
func main() {
  
	database.Connect()
	//http.HandleFunc("/insert", handler.PostData)
	http.HandleFunc("/fetch", handler.GetData)
	http.HandleFunc("/update", handler.UpdateData)
	http.HandleFunc("/delete", handler.DeleteData)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.DeleteData)
	http.ListenAndServe(":8000", nil)
}
