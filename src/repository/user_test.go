package repository

import (
	"fmt"
	"testing"

	"github.com/k197781/go-jwt/src/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	db := db.NewDB().Connect
	dao := NewUserRepository(db)

	email := "test@example.com"
	name := "yuna"
	password := "password"

	t.Run("Cretae", func(t *testing.T) {
		asserts := assert.New(t)

		user, err := dao.CreateUser(email, name, password)
		asserts.NoError(err)
		fmt.Println(user)
		asserts.Equal(name, user.Name)
	})

	t.Run("SignIn", func(t *testing.T) {
		asserts := assert.New(t)

		email := "test@example.com"
		name := "yuna"
		password := "password"
		user, err := dao.SignInUser(email, password)
		asserts.NoError(err)
		fmt.Println(user)
		asserts.Equal(name, user.Name)
	})
}
