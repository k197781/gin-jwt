package main

import (
	"github.com/k197781/go-jwt/src/entity"
	"github.com/k197781/go-jwt/src/infrastructure/db"
)

func main() {
	db := db.NewDB()

	db.Connect.DropTable(entity.User{})

	db.Connect.AutoMigrate(entity.User{})
}
