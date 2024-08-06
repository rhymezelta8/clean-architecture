package user

import (
	"arquitectura/domain/model"
	"github.com/google/uuid"
)

type StorageUser interface {
	GetStorage() (model.Users, error)
	CreateStorage(user model.User) (*uuid.UUID, error)
}
