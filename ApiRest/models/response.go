package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

const contentTypeJSON = "application/json"

func CreateDefaultResponse(rw http.ResponseWriter) *Response {
	return &Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: contentTypeJSON,
	}
}

// respond to the client
func (resp *Response) Send() {
	// Check if the header has been written
	if resp.respWrite.Header().Get("Content-Type") == "" {
		// modify the header
		resp.respWrite.Header().Set("Content-Type", resp.contentType)
		resp.respWrite.WriteHeader(resp.Status)
	}

	// Handle JSON encoding errors
	err := json.NewEncoder(resp.respWrite).Encode(resp.Data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

// changes to return the data to the client
func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

// errors in listing, deleting or obtaining data //method to respond to an error
func (resp *Response) NotFound(message string) {
	resp.Status = http.StatusNotFound
	resp.Message = message
	resp.Send()
}

// respond error to client
func SendNotFound(rw http.ResponseWriter, message string) {
	response := CreateDefaultResponse(rw)
	response.NotFound(message)
	response.Send()
}

// errors entering or updating
func (resp *Response) UnprocessableEntity(message string) {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = message
	resp.Send()
}

func SendUnprocessableEntity(rw http.ResponseWriter, message string) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity(message)
	response.Send()
}

func SendInternalServerError(rw http.ResponseWriter, message string) {
	response := CreateDefaultResponse(rw)
	response.Status = http.StatusInternalServerError
	response.Message = message
	response.Send()
}
