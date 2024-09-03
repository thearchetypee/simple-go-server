package handlers

import (
	"fmt"
	"net/http"
)

func HandleMongoFetch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Data fetched from mongo db")
}
