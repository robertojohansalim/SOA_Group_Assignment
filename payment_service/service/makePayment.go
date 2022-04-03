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

const (
	DEFAULT_USER_ID = "user-id"
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
	CallbackURL    string `json:"callback_url"`
}

type MakePaymentRes struct {
	PaymentID string `json:"payment_id"`
}

func (ths *paymentService) MakePayment(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writeResponse(responseWriter, http.StatusBadRequest, "Unsupported Method")
		return
	}
	// // byts, _ := json.MarshalIndent(, "", " ")
	fmt.Println("\033[36m", string("make Payment Called"), "\033[0m")

	// userID, err := AuthorizeGetUserID(request, ths.paymentModel)
	// if err != nil {
	// 	writeResponse(responseWriter, http.StatusUnauthorized, fmt.Sprintf("Unauthorized Access: %v", err.Error()))
	// 	return
	// }

	var req MakePaymentReq
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(body, &req)
	fmt.Println("\033[36m", string("make Payment Called"), "\033[0m")
	bytsTe, _ := json.MarshalIndent(req, "", " ")
	fmt.Println("\033[36m", string(bytsTe), "\033[0m")
	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	expiredAt := time.Now().Add(time.Duration(req.ActiveDuration) * time.Second).Unix()
	record, err := ths.paymentModel.InsertPaymentRecord(model.PaymentRecordModel{
		UserID:     DEFAULT_USER_ID,
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
		PaymentID: record.ID,
	}
	byts, _ := json.Marshal(res)
	writeResponse(responseWriter, http.StatusOK, string(byts))
	return
}
