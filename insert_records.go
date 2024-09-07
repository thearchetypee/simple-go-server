package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/simple-go-server/db"
	"github.com/simple-go-server/models"
)

// Call this method to insert data into db
func insert() {
	uri := os.Getenv("MONGO_URI")

	client, err := db.ConnectToMongoDb(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	// Get a handle for your collection
	collection := client.Database("simple-go-server").Collection("records")

	// Read JSON file
	file, err := ioutil.ReadFile("records.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON data
	var records []models.Record
	err = json.Unmarshal(file, &records)
	if err != nil {
		log.Fatal(err)
	}

	// Convert records to interface slice for bulk insert
	var interfaces []interface{}
	for _, record := range records {
		interfaces = append(interfaces, record)
	}

	// Insert records
	insertResult, err := collection.InsertMany(context.TODO(), interfaces)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %v documents into MongoDB!\n", len(insertResult.InsertedIDs))
}
