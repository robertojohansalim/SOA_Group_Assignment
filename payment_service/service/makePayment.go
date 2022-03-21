package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"robertojohansalim.github.com/payment/model"
)

const MAKE_PAYMENT_PATH = "/api/make_payment"

const (
	UNPAID_STATUS   = "UNPAID"
	PAID_STATUS     = "PAID"
	REJECTED_STATUS = "REJECTED"
)

// Request Example
// {
//     "external_id": "your-unique-id",
//     "method": "BCA_VA",
//     "amount": 20000,
//     "active_duration": 3600 // InSeconds
// }

type MakePaymentReq struct {
	ExternalID     string `json:"external_id"`
	Method         string `json:"method"`
	Amount         int64  `json:"amount"`
	ActiveDuration int64  `json:"active_duration"`
}

type MakePaymentRes struct {
	PaymentLink string `json:"paymentLink"`
}

func (ths *paymentService) MakePayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writeResponse(responseWriter, http.StatusBadRequest, "Unsupported Method")
		return
	}

	userID, err := AuthorizeGetUserID(request, ths.paymentModel)
	if err != nil {
		writeResponse(responseWriter, http.StatusUnauthorized, err.Error())
		return
	}

	var req MakePaymentReq
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

	if !isMethodAvailable(req.Method) {
		writeResponse(responseWriter, http.StatusBadRequest, "unsupported payment method")
		return
	}

	expiredAt := time.Now().Add(time.Duration(req.ActiveDuration) * time.Second).Unix()
	record, err := ths.paymentModel.InsertPaymentRecord(model.PaymentRecordModel{
		UserID:     userID,
		ExternalID: req.ExternalID,
		Method:     req.Method,
		Amount:     req.Amount,
		Status:     UNPAID_STATUS,
		ExpiredAt:  expiredAt,
	})
	if err != nil {
		writeResponse(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	publicURL := ths.publicURL
	if strings.HasPrefix(publicURL, "http://127.0.0.1") {
		publicURL = strings.Replace(publicURL, "http://127.0.0.1", "localhost", 1)
	}
	res := MakePaymentRes{
		PaymentLink: fmt.Sprintf("%s/api/manage_payment/%s/QUERY", publicURL, record.ID),
	}
	byts, _ := json.Marshal(res)
	writeResponse(responseWriter, http.StatusOK, string(byts))
	return
}
