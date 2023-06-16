package entity

import (
	"github.com/into-the-v0id/user-api.go/value_object"
	"time"
)

type User struct {
	Id          value_object.Identifier
	Name        string
	Password    value_object.PasswordHash
	DateCreated time.Time
	DateUpdated time.Time
}

func NewUser(name string, rawPassword string) *User {
	now := time.Now().UTC()

	return &User{
		value_object.NewIdentifier(),
		name,
		value_object.NewPasswordHash(rawPassword),
		now,
		now,
	}
}
