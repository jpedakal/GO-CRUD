package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect func must be Exported, Capitalized, and comment added.
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
}

// InsertData func must be Exported, Capitalized, and comment added.
func InsertData() {

}

// GetData func must be Exported, Capitalized, and comment added.
func GetData() {

}

// UpdateData func must be Exported, Capitalized, and comment added.
func UpdateData() {

}

// DeleteData func must be Exported, Capitalized, and comment added.
func DeleteData() {

}
