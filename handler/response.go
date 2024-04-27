package handler

import (
	"encoding/json"
	"net/http"
)

const (
	messageTypeSuccess = "success"
	messageTypeError   = "error"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(messageType, message string, data interface{}) response {
	return response{
		MessageType: messageType,
		Message:     message,
		Data:        data,
	}
}

func responseJSON(w http.ResponseWriter, status int, messageType, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(newResponse(messageType, message, data))
}
