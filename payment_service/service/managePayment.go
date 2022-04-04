package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const MANAGE_PAYMENT_PATH = "/api/manage_payment"

type ManagePaymentReq struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Action     string `json:"action"`
}
type ManagePaymentRes struct {
}

func (ths *paymentService) ManagePayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writeResponse(responseWriter, http.StatusBadRequest, "Unsupported Method")
		return
	}

	var req ManagePaymentReq
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	paymentRecordID := req.ID
	paymentStatus := ""
	switch req.Action {
	case "PAY":
		paymentStatus = PAID_STATUS
	case "CANCEL":
		paymentStatus = REJECTED_STATUS
	}

	if paymentRecordID == "" || paymentStatus == "" {
		writeResponse(responseWriter, http.StatusNotFound, "")
		return
	}

	paymentRecord, err := ths.paymentModel.GetPaymentRecordByID(paymentRecordID)
	if err != nil {
		log.Fatal(err)
		writeResponse(responseWriter, http.StatusNotFound, "")
		return
	}

	if paymentRecord.Status != UNPAID_STATUS {
		// writeResponse(responseWriter, http.StatusTemporaryRedirect, fmt.Sprintf("/api/manage_payment/%s/QUERY", paymentRecordID))
		writeResponse(responseWriter, http.StatusAlreadyReported, "STATUS CANNOT BE CHANGED")
		return
	}

	ths.paymentModel.UpdatePaymentStatusRecordByID(paymentRecordID, paymentStatus)
	var response string
	writeResponse(responseWriter, http.StatusOK, response)
	return
}

const (
	PAY_EVENT    = "PAY_EVENT"
	REJECT_EVENT = "REJECT_EVENT"
)

type PaymentCallbackEvent struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
	Method     string `json:"method"`
	Amount     int64  `json:"amount"`
	ExpiryDate string `json:"expiry_date"`
}
