package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PaymentDatabaseModel interface {
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
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.DatabaseName)

	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &paymentDatabase{
		PaymentDatabaseModelConfig: config,
		DB:                         DB,
	}
}

type paymentDatabase struct {
	PaymentDatabaseModelConfig
	DB *sql.DB
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

func InsertPaymentRecord() {

}

func UpdatePaymentRecord() {

}
