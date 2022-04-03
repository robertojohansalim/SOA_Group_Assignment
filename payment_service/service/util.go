package service

import (
	"fmt"
	"net/http"
	"os"

	"robertojohansalim.github.com/payment/model"
)

func AuthorizeGetUserIDWithCallback(request *http.Request, databaseModel model.PaymentDatabaseModel) (userID, callbackURL string, err error) {
	register_user_id := request.Header.Get("authorization")
	if os.Getenv("WITH_AUTHORIZATION_VALIDATION") == "FALSE" {
		return register_user_id, "", nil
	}

	if register_user_id == "" {
		err = fmt.Errorf("unauthorized access, no authorization header provided")
		return
	}

	userCallback, err := databaseModel.GetUserCallback(register_user_id)
	if err != nil {
		return
	}
	if userCallback.ID == "" {
		err = fmt.Errorf("unauthorized access, userId '%s' have not been registered", register_user_id)
		return
	}

	return userCallback.UserID, userCallback.CallbackURL, nil
}

func AuthorizeGetUserID(request *http.Request, databaseModel model.PaymentDatabaseModel) (userID string, err error) {
	userID, _, err = AuthorizeGetUserIDWithCallback(request, databaseModel)
	return userID, err
}

func writeResponse(responseWriter http.ResponseWriter, StatusCode int, message string) {
	// Bypass CORS
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	responseWriter.WriteHeader(StatusCode)
	responseWriter.Write([]byte(message))
}

var mapAvailableMethod = map[string]bool{
	"BANK_TRANSFER": true,
	"BCA_VA":        true,
	"BRI_VA":        true,
	"GOPAY":         true,
	"OVO":           true,
	"QRIS":          true,
}

func isMethodAvailable(method string) bool {
	if isAvailable, ok := mapAvailableMethod[method]; ok {
		return isAvailable
	}
	return false
}
