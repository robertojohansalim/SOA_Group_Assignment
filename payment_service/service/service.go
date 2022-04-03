package service

import (
	"net/http"

	"robertojohansalim.github.com/payment/model"
)

type PaymentService interface {
	MakePayment(http.ResponseWriter, *http.Request)
	GetPayment(http.ResponseWriter, *http.Request)
	ManagePayment(http.ResponseWriter, *http.Request)
}

func NewPaymentService(paymentModel model.PaymentDatabaseModel, publicURL string) PaymentService {
	return &paymentService{
		paymentModel: paymentModel,
		publicURL:    publicURL,
	}
}

type paymentService struct {
	paymentModel model.PaymentDatabaseModel
	publicURL    string
}
