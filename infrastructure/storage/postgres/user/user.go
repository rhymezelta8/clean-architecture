package user

import (
	"arquitectura/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return User{db}
}

func (u User) GetStorage() (model.Users, error) {
	var users model.Users
	u.db.Find(&users)

	return users, nil
}

func (u User) CreateStorage(user model.User) (*uuid.UUID, error) {
	u.db.Create(&user)
	return &user.ID, nil
}
