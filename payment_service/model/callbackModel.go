package model

import (
	"fmt"

	_ "github.com/lib/pq"
)

type UserCallback struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	CallbackURL string `json:"callback_url"`
}

func (ths *paymentDatabase) UpsertUserCallback(spec UserCallback) (UserCallback, error) {
	var id string
	queryString := fmt.Sprintf(
		`INSERT INTO "payment".user_callbacks (user_id, callback_url) VALUES ('%s', '%s')
			ON CONFLICT (user_id)
			DO UPDATE SET callback_url = '%s'`,
		spec.UserID,
		spec.CallbackURL,
		spec.CallbackURL,
	)
	// byts, _ := json.MarshalIndent(, "", " ")
	fmt.Println("\033[36m", queryString, "\033[0m")
	err := ths.db.QueryRow(queryString).Scan(&id)
	if err != nil {
		return UserCallback{}, err
	}

	return ths.GetUserCallback(spec.UserID)
}

func (ths *paymentDatabase) GetUserCallback(userID string) (UserCallback, error) {
	var rowID, callbackURL string
	queryString := fmt.Sprintf(`SELECT id, user_id, callback_url FROM  "payment".user_callbacks WHERE user_id = '%s'`, userID)
	err := ths.db.QueryRow(queryString).Scan(&rowID, &userID, &callbackURL)
	if err != nil {
		return UserCallback{}, err
	}

	return UserCallback{
		ID:          rowID,
		UserID:      userID,
		CallbackURL: callbackURL,
	}, nil
}
