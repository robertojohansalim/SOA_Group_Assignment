package service

import (
	"fmt"
	"net/http"
)

func (ths *paymentService) GetPayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Unsupported Method"))
		return
	}
	fmt.Fprintln(responseWriter, "GetPayment!")
	return
}
