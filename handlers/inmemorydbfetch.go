package handlers

import (
	"fmt"
	"net/http"
)

func HandleFetchFromInMemory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Data fetched from In memory")
}
