package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const GET_PAYMENT_PATH = "/api/get_payment/{id}"

type GetPaymentRes struct {
	ExternalID string `json:"external_id"`
	Method     string `json:"method"`
	Status     string `json:"status"`
	ExpiredAt  string `json:"expired_at"`
	CreatedAt  string `json:"created_at"`
}

func (ths *paymentService) GetPayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Unsupported Method"))
		return
	}

	vars := mux.Vars(request)
	externalID := vars["id"]
	if externalID == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("external_id not found"))
		return
	}

	paymentRecord := ths.paymentModel.GetPaymentRecordByExternalID(externalID)
	res := GetPaymentRes{
		ExternalID: paymentRecord.ExternalID,
		Method:     paymentRecord.Method,
		Status:     paymentRecord.Status,
		ExpiredAt:  time.Unix(paymentRecord.ExpiredAt, 0).Format("2006-02-04 15:04:05"),
		CreatedAt:  time.Unix(paymentRecord.CreatedAt, 0).Format("2006-02-04 15:04:05"),
	}
	byts, err := json.Marshal(res)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseWriter.Write(byts)
	return
}
