package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type paymentDatabase struct {
	PaymentDatabaseModelConfig
	db *sql.DB
}

type PaymentRecordModel struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	ExternalID string `json:"external_id"`
	Amount     int64  `json:"amount"`
	Method     string `json:"method"`
	Status     string `json:"status"`
	ExpiredAt  int64  `json:"expired_at"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (ths *paymentDatabase) InsertPaymentRecord(input PaymentRecordModel) (PaymentRecordModel, error) {
	expiretAtTime := time.Unix(input.ExpiredAt, 0).Format("2006-01-02 15:04:05")
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	queryString := fmt.Sprintf(
		`INSERT INTO "payment".payment_record (user_id, external_id, method, status, amount, expired_at, created_at) VALUES ('%s', '%s', '%s', '%s', '%d', '%s', '%s') ON CONFLICT ON CONSTRAINT payment_record_un DO NOTHING`,
		input.UserID,
		input.ExternalID,
		input.Method,
		input.Status,
		input.Amount,
		expiretAtTime,
		nowTime,
	)
	fmt.Println("\033[36m", string(queryString), "\033[0m")
	_, err := ths.db.Exec(queryString)
	if err != nil {
		return PaymentRecordModel{}, err
	}

	paymentRecord, err := ths.GetPaymentRecordByExternalID(input.UserID, input.ExternalID)
	return paymentRecord, err
}

func (ths *paymentDatabase) GetPaymentRecordByExternalID(query_userId, query_external_id string) (PaymentRecordModel, error) {
	var id, external_id, method, status, userId sql.NullString
	var amount sql.NullInt64
	var expired_at, created_at, updated_at sql.NullTime
	queryString := fmt.Sprintf(
		`SELECT 
			id, 
			user_id,
			external_id, 
			amount,
			method,
			status, 
			expired_at, 
			created_at,
			updated_at 
		FROM  "payment".payment_record 
		WHERE 
			user_id = '%s' AND external_id = '%s'`,
		query_userId,
		query_external_id,
	)
	fmt.Println("\033[36m", string(queryString), "\033[0m")
	err := ths.db.QueryRow(queryString).Scan(&id, &userId, &external_id, &amount, &method, &status, &expired_at, &created_at, &updated_at)
	if err != nil {
		return PaymentRecordModel{}, err
	}

	return PaymentRecordModel{
		ID:         id.String,
		UserID:     userId.String,
		ExternalID: external_id.String,
		Amount:     amount.Int64,
		Method:     method.String,
		Status:     status.String,
		ExpiredAt:  expired_at.Time.Unix(),
		CreatedAt:  created_at.Time.Unix(),
		UpdatedAt:  updated_at.Time.Unix(),
	}, nil
}

func (ths *paymentDatabase) GetPaymentRecordByID(query_id string) (PaymentRecordModel, error) {
	var id, external_id, method, status, userId sql.NullString
	var amount sql.NullInt64
	var expired_at, created_at, updated_at sql.NullTime
	queryString := fmt.Sprintf(
		`SELECT 
			id, 
			user_id,
			external_id, 
			amount,
			method,
			status, 
			expired_at, 
			created_at,
			updated_at 
		FROM  "payment".payment_record 
		WHERE 
			id = '%s'`,
		query_id,
	)
	fmt.Println("\033[36m", string(queryString), "\033[0m")
	err := ths.db.QueryRow(queryString).Scan(&id, &userId, &external_id, &amount, &method, &status, &expired_at, &created_at, &updated_at)
	if err != nil {
		return PaymentRecordModel{}, err
	}

	return PaymentRecordModel{
		ID:         id.String,
		UserID:     userId.String,
		ExternalID: external_id.String,
		Amount:     amount.Int64,
		Method:     method.String,
		Status:     status.String,
		ExpiredAt:  expired_at.Time.Unix(),
		CreatedAt:  created_at.Time.Unix(),
		UpdatedAt:  updated_at.Time.Unix(),
	}, nil
}

func (ths *paymentDatabase) UpdatePaymentStatusRecordByID(recordId, newStatus string) (PaymentRecordModel, error) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	queryString := fmt.Sprintf(
		`UPDATE "payment".payment_record
		SET status = '%s',
		updated_at = '%s'
		WHERE id = '%s';`,
		newStatus,
		nowTime,
		recordId,
	)
	// byts, _ := json.MarshalIndent(, "", " ")
	fmt.Println("\033[36m", string(queryString), "\033[0m")
	err := ths.db.QueryRow(queryString).Scan()
	if err != nil {
		return PaymentRecordModel{}, err
	}

	paymentRecord, err := ths.GetPaymentRecordByID(recordId)
	return paymentRecord, err
}

func (ths *paymentDatabase) UpdatePaymentStatusRecord(userId, external_id, newStatus string) (PaymentRecordModel, error) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	queryString := fmt.Sprintf(
		`UPDATE "payment".payment_record
		SET status = '%s',
		updated_at = '%s'
		WHERE user_id = '%s' AND external_id = '%s';`,
		newStatus,
		nowTime,
		userId,
		external_id,
	)
	err := ths.db.QueryRow(queryString).Scan()
	if err != nil {
		return PaymentRecordModel{}, err
	}

	paymentRecord, err := ths.GetPaymentRecordByExternalID(userId, external_id)
	return paymentRecord, err
}
