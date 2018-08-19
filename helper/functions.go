package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response struct
type Response struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

// RespondWithError prepares error JSON message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	serveResponse := &Response{Code: code, Status: false, Data: nil, Error: msg}
	response, _ := json.Marshal(serveResponse)
	fmt.Println()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(string(response)))
}

// RespondWithJSON prepares JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	serveResponse := &Response{Code: code, Status: true, Data: payload, Error: nil}
	response, _ := json.Marshal(serveResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(string(response)))
}
