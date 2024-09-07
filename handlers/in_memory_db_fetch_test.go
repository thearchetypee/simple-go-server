package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simple-go-server/models"
)

func TestHandleFetchFromInMemory(t *testing.T) {
	// Clear the in-memory database before each test
	inMemoryDb = make(map[string]string)

	t.Run("POST - Success", func(t *testing.T) {
		payload := models.InMemoryResponse{Key: "testKey", Value: "testValue"}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/in-memory", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		err := HandleFetchFromInMemory(w, req)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		var response models.InMemorySuccessReponse
		json.Unmarshal(w.Body.Bytes(), &response)
		if response.Msg != "Success" {
			t.Errorf("Expected message 'Success', got %s", response.Msg)
		}
	})

	t.Run("POST - Empty Key", func(t *testing.T) {
		payload := models.InMemoryResponse{Key: "", Value: "testValue"}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/in-memory", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		err := HandleFetchFromInMemory(w, req)

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}
		clientErr, ok := err.(*ClientError)
		if !ok {
			t.Fatalf("Expected ClientError, got %T", err)
		}
		if clientErr.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, clientErr.StatusCode)
		}
	})

	t.Run("GET - Success", func(t *testing.T) {
		// First, add a key-value pair to the in-memory database
		inMemoryDb["testKey"] = "testValue"

		req := httptest.NewRequest(http.MethodGet, "/in-memory?key=testKey", nil)
		w := httptest.NewRecorder()

		err := HandleFetchFromInMemory(w, req)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		var response models.InMemoryResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		if response.Key != "testKey" || response.Value != "testValue" {
			t.Errorf("Expected {testKey, testValue}, got {%s, %s}", response.Key, response.Value)
		}
	})

	t.Run("GET - Key Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/in-memory?key=nonexistentKey", nil)
		w := httptest.NewRecorder()

		err := HandleFetchFromInMemory(w, req)

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}
		if err.Error() != "key not found in db" {
			t.Errorf("Expected 'key not found in db' error, got %v", err)
		}
	})

	t.Run("GET - Missing Key Parameter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/in-memory", nil)
		w := httptest.NewRecorder()

		err := HandleFetchFromInMemory(w, req)

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}
		clientErr, ok := err.(*ClientError)
		if !ok {
			t.Fatalf("Expected ClientError, got %T", err)
		}
		if clientErr.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, clientErr.StatusCode)
		}
	})

	t.Run("Unsupported Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/in-memory", nil)
		w := httptest.NewRecorder()

		err := HandleFetchFromInMemory(w, req)

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}
		clientErr, ok := err.(*ClientError)
		if !ok {
			t.Fatalf("Expected ClientError, got %T", err)
		}
		if clientErr.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, clientErr.StatusCode)
		}
	})
}
