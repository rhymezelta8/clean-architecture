package handler

import (
	"arquitectura/domain/model"
	routerUser "arquitectura/infrastructure/handler/user"
)

func InitRoutes(rs model.RouterSpecification) {
	// User
	routerUser.NewRouter(rs)
}
