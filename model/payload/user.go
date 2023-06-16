package payload

import (
	"github.com/into-the-v0id/user-api.go/value_object"
	"time"
)

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Id          value_object.Identifier `json:"id"`
	Name        string                  `json:"name"`
	DateCreated time.Time               `json:"dateCreated"`
	DateUpdated time.Time               `json:"dateUpdated"`
}
