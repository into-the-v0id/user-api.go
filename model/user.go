package model

import (
	"github.com/into-the-v0id/user-api.go/entity"
	"github.com/into-the-v0id/user-api.go/value_object"
	"time"
)

type PublicUser struct {
	Id          value_object.Identifier `json:"id"`
	Name        string                  `json:"name"`
	DateCreated time.Time               `json:"dateCreated"`
	DateUpdated time.Time               `json:"dateUpdated"`
}

func NewPublicUser(user *entity.User) *PublicUser {
	return &PublicUser{
		user.Id,
		user.Name,
		user.DateCreated,
		user.DateUpdated,
	}
}
