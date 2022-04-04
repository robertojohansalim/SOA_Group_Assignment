package model

import (
	"database/sql"
	"fmt"
	"log"
)

type PaymentDatabaseModel interface {
	InsertPaymentRecord(PaymentRecordModel) (PaymentRecordModel, error)
	GetPaymentRecordByID(query_id string) (PaymentRecordModel, error)
	GetPaymentRecordByExternalID(query_userId, query_external_id string) (PaymentRecordModel, error)
	UpdatePaymentStatusRecordByID(string, string) (PaymentRecordModel, error)
	UpdatePaymentStatusRecord(string, string, string) (PaymentRecordModel, error)
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
	useInMemoryOnly bool,
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
		useInMemoryOnly:            useInMemoryOnly,
	}
}
