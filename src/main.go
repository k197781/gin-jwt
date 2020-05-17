package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k197781/go-jwt/src/handler"
	"github.com/k197781/go-jwt/src/infrastructure/db"
	"github.com/k197781/go-jwt/src/repository"
	"github.com/k197781/go-jwt/src/usecase"
)

func main() {
	r := gin.Default()

	db := db.NewDB().Connect
	userRepository := repository.NewUserRepository(db)
	userInteractor := usecase.NewUserInteractor(userRepository)
	userHandler := handler.NewUserHandler(userInteractor)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.GET("/user", userHandler.GetUserById)

	r.Run()
}
