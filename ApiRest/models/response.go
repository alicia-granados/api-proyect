package models

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response represents the structure for API responses
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

const contentTypeJSON = "application/json"

// CreateDefaultResponse creates a default response object with OK status
func CreateDefaultResponse(rw http.ResponseWriter) *Response {
	return &Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: contentTypeJSON,
	}
}

// HandleError handles errors centrally
func HandleError(rw http.ResponseWriter, status int, message string, err error) {
	log.Printf("Error: %v\n", err)
	response := CreateDefaultResponse(rw)
	response.Status = status
	response.Message = message
	response.Send()
}

// SendData responds with successful data
func SendData(rw http.ResponseWriter, data interface{}, message string, status int) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Message = message
	response.Status = status
	response.Send()
}

// Send responds to the client
func (resp *Response) Send() {
	// Check if the header has been written
	if resp.respWrite.Header().Get("Content-Type") == "" {
		// Modify the header
		resp.respWrite.Header().Set("Content-Type", resp.contentType)
		resp.respWrite.WriteHeader(resp.Status)
	}

	// Handle JSON encoding errors
	if err := json.NewEncoder(resp.respWrite).Encode(resp.Data); err != nil {
		HandleError(resp.respWrite, http.StatusInternalServerError, "Error encoding JSON", err)
	}
}
