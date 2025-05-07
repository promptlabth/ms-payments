package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/controllers"
	"gorm.io/gorm"
)

func UserMessageRoute(r *gin.Engine, DB *gorm.DB) {

	UserMessageController := controllers.NewUserMessageController()

	r.POST("/user-message", func(c *gin.Context) {
		UserMessageController.ResetMessage(DB)
	})
}
