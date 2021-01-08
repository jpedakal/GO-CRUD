package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	dbConn= client.Database("godata")
}

// InsertData to database
func InsertData(name string, age int, city string) {
	
	collection := dbConn.Collection("users")

	result, err := collection.InsertOne(context.TODO(),)
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
