package api

import (
	"encoding/json"
	"io"
	"net/http"
	database "utils"
)

type myData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
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
	database.InsertData(data.Name, data.Age, data.City)

}

// HelloHandler to send data
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome To 2021\n")
}
