package user

import (
	useCaseUser "arquitectura/application/usecase/user"
	"arquitectura/domain/model"
	storageUser "arquitectura/infrastructure/storage/postgres/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)
	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseUser.New(storageUser.New(specification.DB))
	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/users", middlewares...)
	routes.GET("/", h.get)
	routes.POST("/user", h.Create)

}
