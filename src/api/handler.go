package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type myData struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	City string   `json:"city"`
}

// GetData from database
func GetData() {

}

// PostData into database
func PostData(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("data", data)
}

// HelloHandler to send data
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome To 2021\n")
}