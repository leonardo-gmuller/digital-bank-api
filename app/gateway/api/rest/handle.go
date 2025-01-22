package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJSON(rw http.ResponseWriter, statusCode int, payload any, header map[string]string) error {
	for key, value := range header {
		rw.Header().Set(key, value)
	}

	if payload == nil {
		rw.WriteHeader(statusCode)

		return nil
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	err := json.NewEncoder(rw).Encode(payload)
	if err != nil {
		return fmt.Errorf("send json encode: %w", err)
	}

	return nil
}
