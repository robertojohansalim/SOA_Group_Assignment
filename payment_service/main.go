package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"robertojohansalim.github.com/payment/model"
	"robertojohansalim.github.com/payment/service"
)

const (
	DEFAULT_PORT = "8080"
	HOST         = "http://127.0.0.1"
)

func main() {
	godotenv.Load()

	// Reading Environment Variable
	port := DEFAULT_PORT
	if customPort := os.Getenv("HTTP_SERVER_PORT"); customPort != "" {
		port = customPort
	}

	paymentModel := model.MakePaymentModel(
		model.PaymentDatabaseModelConfig{
			Host:         "127.0.0.1",
			Port:         "5432",
			User:         "postgres",
			Password:     "mysecret",
			DatabaseName: "paymentService",
		},
		false,
	)

	paymentService := service.NewPaymentService(paymentModel, HOST+":"+port)

	r := mux.NewRouter()
	// Manage Payment Endpoints
	r.HandleFunc(service.MAKE_PAYMENT_PATH, paymentService.MakePayment)
	r.HandleFunc(service.GET_PAYMENT_PATH, paymentService.GetPayment)
	r.HandleFunc(service.MANAGE_PAYMENT_PATH, paymentService.ManagePayment)

	// Register Callback Endpoints
	r.HandleFunc(service.REGISTER_CALLBACK_PATH, paymentService.RegisterCallback)

	// Start Serving Server
	log.Println("Starting Server in port: " + port)
	http.ListenAndServe(":"+port, r)
}
