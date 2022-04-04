package service

import (
	"net/http"
)

func writeResponse(responseWriter http.ResponseWriter, StatusCode int, message string) {
	// Bypass CORS
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	responseWriter.WriteHeader(StatusCode)
	responseWriter.Write([]byte(message))
}

var mapAvailableMethod = map[string]bool{
	"BANK_TRANSFER": true,
	"BCA_VA":        true,
	"BRI_VA":        true,
	"GOPAY":         true,
	"OVO":           true,
	"QRIS":          true,
}

func isMethodAvailable(method string) bool {
	if isAvailable, ok := mapAvailableMethod[method]; ok {
		return isAvailable
	}
	return false
}
