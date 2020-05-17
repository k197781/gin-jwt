package main

import (
	"github.com/k197781/go-jwt/src/entity"
	"github.com/k197781/go-jwt/src/infrastructure/db"
)

func main() {
	db := db.NewDB()

	user := entity.User{
		ID:   0,
		Name: "yuna",
	}
	db.Connect.Create(&user)
}
