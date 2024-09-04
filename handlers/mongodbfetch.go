package handlers

import (
	"fmt"
	"net/http"
)

func HandleMongoFetch(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Data fetched from mongo db")
	return nil
}
