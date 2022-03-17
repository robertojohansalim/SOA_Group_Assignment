package main

import (
	"log"
	"net/http"
	"os"

	"robertojohansalim.github.com/payment/model"
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

	paymentModel := model.MakePaymentModel(model.PaymentDatabaseModelConfig{
		Host:         "127.0.0.1",
		Port:         "5432",
		User:         "postgres",
		Password:     "mysecret",
		DatabaseName: "paymentService",
	})

	paymentService := service.NewPaymentService(paymentModel)

	// Mapping Endpoints
	http.HandleFunc(MAKE_PAYMENT_PATH, paymentService.MakePayment)
	http.HandleFunc(GET_PAYMENT_PATH, paymentService.GetPayment)
	http.HandleFunc(COMPLETE_PAYMENT_PATH, paymentService.CompletePayment)

	// Start Serving Server
	log.Println("Starting Server in port: " + port)
	http.ListenAndServe(":"+port, nil)
}
