package model

import "time"

type User struct {
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name" `
	ApiKey    string    `json:"api_key"`
}
