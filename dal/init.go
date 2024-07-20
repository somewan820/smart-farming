package dal

import (
	"ST/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MysqlDefaultDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
