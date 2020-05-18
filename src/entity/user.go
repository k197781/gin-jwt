package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"-"`
	Password  string     `json:"-"`
	CreatedAt *time.Time `json:"-"`
}
