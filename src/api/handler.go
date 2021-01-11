package api

import (
	"encoding/json"
	"fmt"
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
func PostData(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	res, err := database.InsertData(data.Name, data.Age, data.City)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		fmt.Println(res)
		io.WriteString(w, "Document created")
	}
}