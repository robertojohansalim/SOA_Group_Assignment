package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type PaymentDatabaseModel interface {
	InsertPaymentRecord(PaymentRecordModel)
	GetPaymentRecordByExternalID(string) PaymentRecordModel
	UpdatePaymentRecord()
}

type PaymentDatabaseModelConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

func MakePaymentModel(
	config PaymentDatabaseModelConfig,
) PaymentDatabaseModel {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &paymentDatabase{
		PaymentDatabaseModelConfig: config,
		db:                         db,
	}
}

type paymentDatabase struct {
	PaymentDatabaseModelConfig
	db *sql.DB
}

type PaymentRecordModel struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Method     string `json:"method"`
	Status     string `json:"status"`
	ExpiredAt  int64  `json:"expired_at"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (ths *paymentDatabase) InsertPaymentRecord(input PaymentRecordModel) {
	var id string
	expiretAtTime := time.Unix(input.ExpiredAt, 0).Format("2006-01-02 15:04:05")
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	queryString := fmt.Sprintf(
		`INSERT INTO "payment".payment_record (external_id, method, status, expired_at, created_at) VALUES (%s, %s, %s, %s, %s)`,
		input.ExternalID,
		input.Method,
		input.Status,
		expiretAtTime,
		nowTime,
	)
	ths.db.QueryRow(queryString).Scan(&id)

	ths.GetPaymentRecordByExternalID(input.ExternalID)
}

func (ths *paymentDatabase) GetPaymentRecordByExternalID(query_external_id string) PaymentRecordModel {
	var id, external_id, method, status string
	var expired_at, created_at, updated_at time.Time
	queryString := fmt.Sprintf(`SELECT id, external_id, method, status, expired_at, created_at, updated_at FROM  "payment".payment_record WHERE external_id = '%s'`, query_external_id)
	ths.db.QueryRow(queryString).Scan(&id, &external_id, &method, &status, &expired_at, &created_at, &updated_at)

	return PaymentRecordModel{
		ID:         id,
		ExternalID: external_id,
		Method:     method,
		Status:     status,
		ExpiredAt:  expired_at.Unix(),
		CreatedAt:  created_at.Unix(),
		UpdatedAt:  updated_at.Unix(),
	}
}

func (ths *paymentDatabase) UpdatePaymentRecord() {
	// var id string
	// expiretAtTime := time.Unix(input.ExpiredAt, 0).Format("2014-02-04")
	// nowTime := time.Now().Format("2014-02-04")
	// queryString := fmt.Sprintf(
	// 	`INSERT INTO "payment".payment_record (external_id, method, status, expired_at, created_at) VALUES (%s, %s, %s, %s, %s)`,
	// 	input.ExternalID,
	// 	input.Method,
	// 	input.Status,
	// 	expiretAtTime,
	// 	nowTime,
	// )
	// ths.db.QueryRow(queryString).Scan(&id)

	// ths.GetPaymentRecordByExternalID(input.ExternalID)

}
