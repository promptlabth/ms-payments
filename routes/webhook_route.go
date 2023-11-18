package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WebhookRoute(r *gin.Engine, DB *gorm.DB) {
	r.POST("/webhook/subscription", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Test": c.Request.Body,
		})
	})
}
