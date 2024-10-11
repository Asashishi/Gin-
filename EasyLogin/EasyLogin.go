package main

import (
	"EasyLogin/dataBaseHandler"
	"EasyLogin/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err0 := dataBaseHandler.InitDataBase()
	if err0 != nil {
		log.Fatal(err0)
	}

	/*
		err1 := db.AutoMigrate(&entity.User{})
		if err1 != nil {
			log.Println(err1.Error())
		}
	*/

	// gin框架强制彩色日志
	gin.ForceConsoleColor()
	Server := gin.Default()

	// 注册路由
	routers.LoginRouter(Server, db)

	err2 := Server.Run(":80")
	if err2 != nil {
		log.Fatal(err2)
	}

}
