package value_object

import (
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHash string

func NewPasswordHash(rawPassword string) PasswordHash {
	return PasswordHash(hash(rawPassword))
}

func hash(rawPassword string) (hash string) {
	hashedPassword := lo.Must(bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost))

	return string(hashedPassword)
}

func (passwordHash PasswordHash) Verify(rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(rawPassword))
	if err != nil {
		return false
	}

	return true
}

func (passwordHash PasswordHash) String() string {
	return string(passwordHash)
}
