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
var HelloHandler = func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world\n")
}
