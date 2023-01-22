package db

import (
	"fmt"
	"nokogiri-api/utils"

	"strconv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Db *gorm.DB
)

func InitDB() (err error) {
	port, _ := strconv.Atoi(utils.DbPort)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		utils.ServerName, utils.DbUser, utils.DbPassword, port, utils.DbName)
	Db, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return
	}

	if err = Db.AutoMigrate(&UserInfo{}, &Post{}); err != nil {
		return
	}
	return
}
