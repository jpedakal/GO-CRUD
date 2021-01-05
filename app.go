package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	

	collection := client.Database("test").Collection("trainers")

	jayakrishna := Trainer{"Jayakrishna", 27, "Bangalore"}

	insertResult, err := collection.InsertOne(context.TODO(), jayakrishna)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted document", insertResult.InsertedID)

}
