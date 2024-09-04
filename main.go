package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/simple-go-server/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/mongo", handlers.Make(handlers.HandleMongoFetch))
	http.HandleFunc("/in-memory", handlers.Make(handlers.HandleFetchFromInMemory))

	fmt.Printf("Server is starting on port %s...\n", os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe(os.Getenv("SERVER_PORT"), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
