package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k197781/go-jwt/src/usecase"
)

type UserHandler struct {
	userInteractor usecase.UserInteractor
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	user, err := h.userInteractor.GetUserById(int64(1))
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, user)
	return
}

func NewUserHandler(
	userInteractor usecase.UserInteractor) *UserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
	}
}
