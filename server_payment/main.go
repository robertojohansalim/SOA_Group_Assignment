package main

import (
	"log"
	"net/http"
	"os"

	"robertojohansalim.github.com/payment/service"
)

const (
	MAKE_PAYMENT_PATH     = "/api/make_payment"
	GET_PAYMENT_PATH      = "/api/get_payment"
	COMPLETE_PAYMENT_PATH = "/api/complete_payment"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	// Reading Environment Variable
	port := DEFAULT_PORT
	if customPort := os.Getenv("HTTP_SERVER_PORT"); customPort != "" {
		port = customPort
	}

	paymentService := service.NewPaymentService()

	// Mapping Endpoints
	http.HandleFunc(MAKE_PAYMENT_PATH, paymentService.MakePayment)
	http.HandleFunc(GET_PAYMENT_PATH, paymentService.GetPayment)
	http.HandleFunc(COMPLETE_PAYMENT_PATH, paymentService.CompletePayment)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "halo!")
	// })

	// Start Serving Server
	log.Println("Starting Server in port: " + port)
	http.ListenAndServe(":"+port, nil)
}
