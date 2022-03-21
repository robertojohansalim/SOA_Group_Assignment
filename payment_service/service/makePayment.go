package service

import (
	"fmt"
	"net/http"
)

const MAKE_PAYMENT_PATH = "/api/make_payment"

func (ths *paymentService) MakePayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Unsupported Method"))
		return
	}
	fmt.Fprintln(responseWriter, "MakePayment!")
	return
}
