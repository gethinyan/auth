package models

import (
	"fmt"
	"log"

	"e.coding.net/handnote/handnote/pkg/setting"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql" // mysql driver
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// init 初始化数据库连接
func init() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.Database.User,
		setting.Database.Password,
		setting.Database.Host,
		setting.Database.Port,
		setting.Database.Dbname,
	)
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	dbConn.AutoMigrate(&User{})
	dbConn.AutoMigrate(&Memo{})
	dbConn.AutoMigrate(&Version{})
	dbConnDB, err := dbConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	dbConnDB.SetMaxIdleConns(10)
	dbConnDB.SetMaxOpenConns(100)
}
