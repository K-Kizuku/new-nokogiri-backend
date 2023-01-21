package db

import (
	"fmt"
	"nokogiri-api/utils"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDB() (err error) {
	fmt.Println(utils.DbPassword)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", utils.ServerName, utils.DbUser, utils.DbPassword, utils.DbName, utils.DbPort)
	fmt.Println(dsn)
	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	if err = Db.AutoMigrate(&User{}, &Post{}); err != nil {
		return
	}
	return
}
