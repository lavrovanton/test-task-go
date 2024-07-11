package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"test-task-go/internal/config"
	"test-task-go/internal/db"
	"test-task-go/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const keyLength = 64

func generateString(length int) string {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	return base64.URLEncoding.EncodeToString(buffer)[:length]
}

func main() {
	cfg := config.Get()
	db := db.Get(cfg)

	key := generateString(keyLength)

	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if bcrypt.CompareHashAndPassword(hash, []byte(key)) != nil {
		log.Fatal(err)
	}

	var user model.User

	result := db.Where("name = ?", "admin").First(&user)

	user.Name = "admin"
	user.ApiKey = string(hash)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&user)
	} else {
		user.Name = "admin"
		user.ApiKey = string(hash)
		db.Save(&user)
	}

	fmt.Println(key)
}
