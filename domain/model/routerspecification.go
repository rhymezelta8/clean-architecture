package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterSpecification struct {
	Api *gin.Engine
	DB  *gorm.DB
}
