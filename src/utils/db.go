package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post to export data
type Post struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	City string `json:"city,omitempty"`
}

// InsertOneResult exporting
type InsertOneResult struct {
	InsertedID interface{}
}

var dbConn *mongo.Database

// Connect to database
func Connect() {
	URI := "mongodb://127.0.0.1:27017/?gssapiServiceName=mongodb"
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	dbConn = client.Database("godata")
}

// InsertData to database
func InsertData(name string, age int, city string) (*InsertOneResult, error) {
	post := Post{name, age, city}
	collection := dbConn.Collection("users")

	result, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}

	id := result.InsertedID

	return id, err
}

// GetData from database
func GetData() {

}

// UpdateData in database
func UpdateData() {

}

// DeleteData from database
func DeleteData() {

}
