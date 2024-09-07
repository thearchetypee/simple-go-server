package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/simple-go-server/config"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

type ErrorResponse struct {
	Error string `json:"error"`
}

func WithConfig(cfg *config.Config, h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.ConfigKey, cfg)
		err := h(w, r.WithContext(ctx))
		if err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := http.StatusInternalServerError
	var clientError *ClientError
	if errors.As(err, &clientError) {
		statusCode = clientError.StatusCode
	}

	slog.Error("HTTP API error",
		"err", err.Error(),
		"path", r.URL.Path,
		"status", statusCode,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

type ClientError struct {
	StatusCode int
	Err        error
}

func (e *ClientError) Error() string {
	return e.Err.Error()
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}

	_, err = w.Write(prettyJSON)
	return err
}
