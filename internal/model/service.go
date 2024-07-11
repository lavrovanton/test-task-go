package model

import "time"

type Service struct {
	Id          uint64    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	PaymentType string    `json:"paymentType"`
	Price       uint64    `json:"price"`
}
