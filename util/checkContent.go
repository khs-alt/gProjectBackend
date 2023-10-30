package util

import (
	"errors"
	"io"
	"net/http"
)

func ProcessRequest(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return nil, errors.New("Invalid Content-Type")
	}

	// 요청 본문 읽기
	body, err := readRequestBody(r)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return nil, errors.New("Error reading request body")
	}
	return body, nil
}

func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}
