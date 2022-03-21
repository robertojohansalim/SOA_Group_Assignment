package service

import (
	"fmt"
	"net/http"
)

const MANAGE_PAYMENT_PATH = "/api/manage_payment"

func (ths *paymentService) ManagePayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Unsupported Method"))
		return
	}
	fmt.Fprintln(responseWriter, "CompletePayment!")
	return

}
