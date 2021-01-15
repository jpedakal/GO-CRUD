package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func InsertData(payload map[string]string) (interface{}, error) {
	//post := Post{payload}
	collection := dbConn.Collection("users")

	result, err := collection.InsertOne(context.TODO(), payload)

	if err != nil {
		log.Fatal(err)
	}

	id := result.InsertedID
	//oid := result.InsertedID.(primitive.ObjectID)
	//slice := oid[:]

	return id, err
}

// GetData from database
func GetData() (interface{}, error) {
	var results []bson.M
	collection := dbConn.Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

// UpdateData in database
func UpdateData(name string) string {
	collection := dbConn.Collection("users")
	opts := options.Update().SetUpsert(true)

	filter := bson.D{primitive.E{Key: "name", Value: name}}
	update := bson.D{primitive.E{Key:"$set",Value: bson.D{primitive.E{Key: "city", Value: "Hyderabad"}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		return "Successfully updated"
	}
	return "Failed to update document"
}

// DeleteData from database
func DeleteData(name string) map[string]string {
	result := map[string]string{"message": "Deleted document successfully"}
	collection := dbConn.Collection("users")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: name}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return result
}

