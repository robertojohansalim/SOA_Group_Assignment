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

func NewPaymentService(paymentModel model.PaymentDatabaseModel) PaymentService {
	return &paymentService{
		paymentModel: paymentModel,
	}
}

type paymentService struct {
	paymentModel model.PaymentDatabaseModel
}
