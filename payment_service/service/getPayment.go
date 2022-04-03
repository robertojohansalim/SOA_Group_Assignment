package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const GET_PAYMENT_PATH = "/api/get_payment/{id}"

type GetPaymentRes struct {
	ExternalID string `json:"external_id"`
	Method     string `json:"method"`
	Status     string `json:"status"`
	Amount     int64  `json:"amount"`
	ExpiredAt  string `json:"expired_at"`
	CreatedAt  string `json:"created_at"`
}

func (ths *paymentService) GetPayment(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("\033[36m", string("GetPayment Called"), "\033[0m")

	if request.Method != "GET" {
		writeResponse(responseWriter, http.StatusBadRequest, "Unsupported Method")
		return
	}

	vars := mux.Vars(request)
	paymentId := vars["id"]
	if paymentId == "" {
		writeResponse(responseWriter, http.StatusBadRequest, "external_id not found")
		return
	}

	paymentRecord, err := ths.paymentModel.GetPaymentRecordByID(paymentId)
	if err != nil {
		writeResponse(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	res := GetPaymentRes{
		ExternalID: paymentRecord.ExternalID,
		Method:     paymentRecord.Method,
		Status:     paymentRecord.Status,
		Amount:     paymentRecord.Amount,
		ExpiredAt:  time.Unix(paymentRecord.ExpiredAt, 0).Format("2006-02-04 15:04:05"),
		CreatedAt:  time.Unix(paymentRecord.CreatedAt, 0).Format("2006-02-04 15:04:05"),
	}
	byts, err := json.Marshal(res)
	if err != nil {
		writeResponse(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(responseWriter, http.StatusOK, string(byts))
	return
}
