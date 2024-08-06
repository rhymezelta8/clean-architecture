package user

import (
	"arquitectura/domain/model"
	"github.com/google/uuid"
)

type UseCaseUser interface {
	Get() (model.Users, error)
	Create(model.User) (*uuid.UUID, error)
}

//go:generate mockery --case=snake --outpkg=usecasemocks --output=../usecasemocks --name=UseCaseUser
