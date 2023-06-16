package repository

import (
	"fmt"
	"github.com/into-the-v0id/user-api.go/entity"
	"github.com/into-the-v0id/user-api.go/value_object"
	"github.com/samber/lo"
)

type UserRepository interface {
	GetById(id value_object.Identifier) *entity.User
	GetAll() []*entity.User
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
}

var User UserRepository = &InMemoryUserRepository{
	make(map[value_object.Identifier]*entity.User),
}

func init() {
	// Create sample user
	newUser := entity.NewUser("max", "hunter2")
	lo.Must0(User.Create(newUser))
}

type InMemoryUserRepository struct {
	users map[value_object.Identifier]*entity.User
}

func (repo *InMemoryUserRepository) GetAll() []*entity.User {
	userList := make([]*entity.User, 0, len(repo.users))
	for _, user := range repo.users {
		userList = append(userList, user)
	}

	return userList
}

func (repo *InMemoryUserRepository) GetById(id value_object.Identifier) *entity.User {
	return repo.users[id]
}

func (repo *InMemoryUserRepository) Create(user *entity.User) error {
	if repo.GetById(user.Id) != nil {
		return fmt.Errorf("user with id %s already exists", user.Id)
	}

	repo.users[user.Id] = user
	return nil
}

func (repo *InMemoryUserRepository) Update(user *entity.User) error {
	if repo.GetById(user.Id) == nil {
		return fmt.Errorf("user with id %s does not exist", user.Id)
	}

	repo.users[user.Id] = user
	return nil
}

func (repo *InMemoryUserRepository) Delete(user *entity.User) error {
	if repo.GetById(user.Id) == nil {
		return fmt.Errorf("user with id %s does not exist", user.Id)
	}

	delete(repo.users, user.Id)
	return nil
}
