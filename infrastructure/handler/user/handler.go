package user

import (
	"arquitectura/domain/model"
	"arquitectura/domain/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	useCase user.UseCaseUser
}

func newHandler(useCase user.UseCaseUser) handler {
	return handler{useCase}
}

func (h handler) get(c *gin.Context) {
	ms, err := h.useCase.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		//c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}
	c.JSON(http.StatusOK, ms)
}

func (h handler) Create(c *gin.Context) {
	var req model.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ms, err := h.useCase.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		//c.JSON(response.Wrong(model.ResponseError{Error: err.Error()}))
		return
	}
	c.JSON(http.StatusCreated, &ms)
}
