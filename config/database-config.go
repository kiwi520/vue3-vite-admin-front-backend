package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang_api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// 初始化数据库
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("获取env文件失败")
	}
	dbUser := os.Getenv("DB_USER")
	dbPASS := os.Getenv("DB_PASS")
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPASS, dbHOST, dbPORT, dbNAME)
	//dsn :=  "root:root@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接mysql数据库失败！")
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
		&entity.Department{},
		)

	return db
}

//关闭数据库连接
func CloseDatabaseConnection(db *gorm.DB) {
	dbHandle, err := db.DB()
	if err != nil {
		panic("连接数据库失败")
	}

	if err := dbHandle.Close(); err != nil {
		panic("关闭数据库失败")
	}
}
