package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"robertojohansalim.github.com/payment/model"
)

const REGISTER_CALLBACK_PATH = "/api/register_callback"

// {
//     "callback_type": "update_status_payment",
//     "callback_url": "http://your.domain/recieve_callback/url"
// }
type RegisterCallbackReq struct {
	CallbackURL string `json:"callback_url"`
}

func (ths *paymentService) RegisterCallback(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writeResponse(responseWriter, http.StatusBadRequest, "Unsupported Method")
		return
	}

	userID, err := AuthorizeGetUserID(request, ths.paymentModel)
	if err != nil {
		writeResponse(responseWriter, http.StatusUnauthorized, err.Error())
		return
	}

	var req RegisterCallbackReq
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

	ths.paymentModel.UpsertUserCallback(model.UserCallback{
		UserID:      userID,
		CallbackURL: req.CallbackURL,
	})

	writeResponse(responseWriter, http.StatusOK, "Successfully Register Callback for "+userID)
	return
}
