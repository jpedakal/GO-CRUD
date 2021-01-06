package api

import (
	"io"
	"net/http"
)

// GetData from database
func GetData() {

}

// PostData into database
func PostData() {

}

// HelloHandler to send data
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome To 2021\n")
}
