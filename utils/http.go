package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func JsonResponse(w http.ResponseWriter, status int, resp interface{}) {
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func DecodeRequest(r *http.Request, req interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(req); err != nil {
		return fmt.Errorf("failed to decode request: %w", err)
	}
	return nil
}

func UrlEscaping(filePath string) (string, error) {
	decodedPathVar, err := url.PathUnescape(filePath)
	if err != nil {
		return "", fmt.Errorf("error decoding the string %w", err)
	}
	return decodedPathVar, nil
}
