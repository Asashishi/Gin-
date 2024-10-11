package dataBaseHandler

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// InitDataBase 初始化数据库连接池
func InitDataBase() (db *gorm.DB, err error) {
	// 定义log对象并输出至控制台
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//// 超时阈值
			//SlowThreshold: time.Second,
			LogLevel: logger.Info,
			// 使用彩色日志
			Colorful: true,
		})
	// 定义数据源
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_base?charset=utf8&parseTime=True&loc=Local"
	db, err0 := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err0 != nil {
		log.Fatal(err0)
	}
	DbPool, err1 := db.DB()
	if err1 != nil {
		log.Fatal(err1)
	}
	// 最大连接数
	DbPool.SetMaxIdleConns(25)
	// 最大空闲连接数
	DbPool.SetMaxOpenConns(15)
	return db, err
}
