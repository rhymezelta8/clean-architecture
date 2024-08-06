package user

import (
	"arquitectura/application/repository/storage/user"
	"arquitectura/domain/model"
	"github.com/google/uuid"
)

type User struct {
	storage user.StorageUser
}

func New(storage user.StorageUser) User {
	return User{storage}
}

func (u User) Get() (model.Users, error) {
	m, err := u.storage.GetStorage()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (u User) Create(user model.User) (*uuid.UUID, error) {
	user.ID = uuid.New()
	m, err := u.storage.CreateStorage(user)
	if err != nil {
		return nil, err
	}

	return m, nil
}
