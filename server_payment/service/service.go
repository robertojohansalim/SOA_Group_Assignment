package service

import "net/http"

type PaymentService interface {
	MakePayment(http.ResponseWriter, *http.Request)
	GetPayment(http.ResponseWriter, *http.Request)
	CompletePayment(http.ResponseWriter, *http.Request)
}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

type paymentService struct {
}
