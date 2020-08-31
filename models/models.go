package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gethinyan/enterprise/pkg/setting"
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
	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

}
