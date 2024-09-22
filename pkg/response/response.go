package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Status200 = http.StatusOK
	Status300 = http.StatusMultipleChoices
	Status400 = http.StatusBadRequest
	Status404 = http.StatusNotFound
	Status500 = http.StatusInternalServerError
	Status504 = http.StatusGatewayTimeout
	Status505 = http.StatusHTTPVersionNotSupported
)

const (
	Message200 = "Request completed successfully"
	Message300 = "Multiple choices"
	Message400 = "Bad request"
	Message404 = "Page not found"
	Message500 = "Internal server error"
)

type response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"message"`
}

func New(code int, body string) []byte {
	response := response{
		StatusCode: code,
		Body:       body,
	}

	json, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Bad marshaling json")
	}

	return json
}
