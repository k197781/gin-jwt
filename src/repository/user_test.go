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

	t.Run("Cretae", func(t *testing.T) {
		asserts := assert.New(t)

		email := "test@example.com"
		name := "yuna"
		password := "password"
		user, err := dao.CreateUser(email, name, password)
		asserts.NoError(err)
		fmt.Println(user)
		asserts.Equal(name, user.Name)
	})
}
