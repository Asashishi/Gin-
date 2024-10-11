package controllers

import (
	"EasyLogin/entity"
	"EasyLogin/utlis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func UserLogin(db *gorm.DB, ctl *gin.Context) {
	var webUserInfo entity.User
	var userInfo entity.User
	err0 := ctl.ShouldBind(&webUserInfo)
	if err0 != nil {
		log.Println("args error", err0)
		APIResponseVo := utlis.Error{Code: 403, Msg: "args error"}
		ctl.JSON(200, APIResponseVo)
		return
	}
	db.Where("username = ?", webUserInfo.Username).First(&userInfo)
	if userInfo.Username == "" || userInfo.Password == "" {
		APIResponseVo := utlis.Error{Code: 403, Msg: "user not exist"}
		ctl.JSON(200, APIResponseVo)
		return
	} else if userInfo.Password != webUserInfo.Password {
		APIResponseVo := utlis.Error{Code: 403, Msg: "password error"}
		ctl.JSON(200, APIResponseVo)
	} else {
		ctl.Set("isLogin", true)
		// session.Set("islogin",true)
		APIResponseVo := utlis.Succeed{Code: 200, Msg: "Succeed", Data: webUserInfo.Username}
		ctl.JSON(200, APIResponseVo)
	}
}
