package routers

import (
	"EasyLogin/routers/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginRouter(Server *gin.Engine, db *gorm.DB) {
	Server.POST("/login", func(ctl *gin.Context) {
		controllers.UserLogin(db, ctl)
	})
}
