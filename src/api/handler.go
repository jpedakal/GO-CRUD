package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	algorithm "secure"
	database "utils"
)

type myData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
}

// GetData from database
func GetData(w http.ResponseWriter, req *http.Request) {
	res, err := database.GetData()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}

// PostData into database
/*
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
*/

// UpdateData of existing documents
func UpdateData(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	res := database.UpdateData(name)

	json.NewEncoder(w).Encode(res)
}

// DeleteData of existing documents
func DeleteData(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	res := database.DeleteData(name)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}

// Register to export
func Register(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var data register
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	text := []byte(data.Password)
	key := []byte("secret")

	hash, err := algorithm.Encrypt(text, key)
	hashPwd := algorithm.BytesToString(hash)
	
	payload := map[string]string{"name": data.Name, "email": data.Email, "password": hashPwd, "city": data.City}
	fmt.Println("hash", hash)
	res, err := database.InsertData(payload)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		fmt.Println(res)
		io.WriteString(w, "Document created")
	}
}

// Login to export
func Login(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	res := database.DeleteData(name)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}
