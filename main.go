package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/simple-go-server/config"
	"github.com/simple-go-server/db"
	"github.com/simple-go-server/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("MONGO_URI")
	client, err := db.ConnectToMongoDb(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	config := config.NewConfig(client)

	http.HandleFunc("/mongo", handlers.WithConfig(config, handlers.HandleMongoFetch))
	http.HandleFunc("/in-memory", handlers.WithConfig(config, handlers.HandleFetchFromInMemory))

	fmt.Printf("Server is starting on port %s...\n", os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe(os.Getenv("SERVER_PORT"), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
