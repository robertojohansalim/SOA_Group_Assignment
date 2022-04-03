package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"robertojohansalim.github.com/payment/model"
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

	var response string
	switch paymentStatus {
	case PAID_STATUS:
		ths.paymentModel.UpdatePaymentStatusRecordByID(paymentRecordID, PAID_STATUS)
	case REJECTED_STATUS:
		notifyCallbackUpdateStatus(ths.paymentModel, paymentRecordID, "REJECT_EVENT")
	default:
		{
			writeResponse(responseWriter, http.StatusNotFound, "")
			return
		}
	}
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

// TODO: Improvement Make this become an EVENT DRIVEN process (using pubsub / maybe golang channel)
// https://medium.com/@ohm.patel1997/publisher-subscriber-architecture-using-golang-5566ca852d9f
func notifyCallbackUpdateStatus(paymentModel model.PaymentDatabaseModel, recordID, EVENT_TITLE string) error {
	paymentRecord, err := paymentModel.GetPaymentRecordByID(recordID)
	if err != nil {
		return err
	}
	paymentCallback, err := paymentModel.GetUserCallback(paymentRecord.UserID)
	if err != nil {
		return err
	}
	var paymentStatus string
	switch EVENT_TITLE {
	case PAY_EVENT:
		paymentStatus = PAID_STATUS
	case REJECT_EVENT:
		paymentStatus = REJECTED_STATUS
	}
	payload := PaymentCallbackEvent{
		ID:         recordID,
		ExternalID: paymentRecord.ExternalID,
		Status:     paymentStatus,
		Method:     paymentRecord.Method,
		Amount:     paymentRecord.Amount,
		ExpiryDate: time.Unix(paymentRecord.ExpiredAt, 0).Format("2006-01-02 15:04:05"),
	}
	byts, _ := json.Marshal(payload)
	_, err = http.Post(paymentCallback.CallbackURL, "application/json", bytes.NewBuffer(byts))
	return err
}

func queryHTMLResponse(recordId string) string {
	response := fmt.Sprintf(
		`
		<!DOCTYPE html>
		<html>
		<body>

		<h1>The button Element</h1>

		<button onclick="location.href = '/api/manage_payment/%s/PAY';" id="myButton" style="font-size:20pt;background-color:#00FF00" >PAY</button>
		<button onclick="location.href = '/api/manage_payment/%s/REJECT';" id="myButton" style="font-size:20pt;background-color:#FF0000" >REJECT</button>

		</body>
		</html>
		`,
		recordId,
		recordId,
	)
	return response
}

func payHTMLResponse() string {
	response :=
		`
		<!DOCTYPE html>
		<html>
		<body>

		<h1>PAYMENT HAS BEEN <b style="color:#00F800">PAID</b></h1>

		
		</body>
		</html>
		`
	return response
}

func rejectHTMLResponse() string {
	response :=
		`
		<!DOCTYPE html>
		<html>
		<body>

		<h1 >PAYMENT HAS BEEN <b style="color:#F80000">REJECTED</b></h1>

		
		</body>
		</html>
		`
	return response
}
