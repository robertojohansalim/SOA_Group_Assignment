package service

import (
	"fmt"
	"net/http"
)

func (ths *paymentService) CompletePayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Unsupported Method"))
		return
	}
	fmt.Fprintln(responseWriter, "CompletePayment!")
	return

}
