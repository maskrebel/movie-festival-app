package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"is_success": true, "message": "Create movie successfully!"})
	}
}
