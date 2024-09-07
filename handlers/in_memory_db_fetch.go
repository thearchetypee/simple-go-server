package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simple-go-server/models"
)

var inMemoryDb = make(map[string]string)

func HandleFetchFromInMemory(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return getInMemoryDb(w, r)
	case http.MethodPost:
		return postInMemoryDb(w, r)
	default:
		return &ClientError{
			StatusCode: http.StatusMethodNotAllowed,
			Err:        fmt.Errorf("method not allowed: %s", r.Method),
		}
	}
}

func getInMemoryDb(w http.ResponseWriter, r *http.Request) error {
	queryParams := r.URL.Query()
	keys, ok := queryParams["key"]
	if !ok || len(keys[0]) < 1 {
		return &ClientError{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("URL Param 'key' is missing"),
		}
	}

	value, ok := inMemoryDb[keys[0]]
	if !ok {
		return fmt.Errorf("key not found in db")
	}

	return writeJSON(w, http.StatusOK, models.InMemoryResponse{Key: keys[0], Value: value})
}

func postInMemoryDb(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body")
	}
	defer r.Body.Close()

	var payload models.InMemoryResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return &ClientError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}
	if len(payload.Key) < 1 {
		return &ClientError{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("key should not be empty"),
		}
	}

	if len(payload.Value) < 1 {
		return &ClientError{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("key should not be empty"),
		}
	}
	inMemoryDb[payload.Key] = payload.Value
	return writeJSON(w, http.StatusOK, &models.InMemorySuccessReponse{Msg: "Success"})
}
