package models

import (
	"api-admin/helpers"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	link *gorm.DB
	err  error
)

func init() {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		helpers.DatabaseConfig.User,
		helpers.DatabaseConfig.Password,
		helpers.DatabaseConfig.Host,
		helpers.DatabaseConfig.Port,
		helpers.DatabaseConfig.Dbname,
		helpers.DatabaseConfig.Charset,
	)

	link, err = gorm.Open(helpers.DatabaseConfig.Type, connectString)

	if err != nil {
		log.Fatal("Fail to connect db error : %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return helpers.DatabaseConfig.TablePrefix + defaultTableName
	}

	link.SingularTable(true)

	link.LogMode(true)
}

func Close() {
	defer link.Close()
}
