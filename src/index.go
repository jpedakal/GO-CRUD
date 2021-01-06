package main

import (
	"fmt"
	mongo "utils"
)

func main() {
	mongo.Connect()
	fmt.Println("Successfully connected")
}
