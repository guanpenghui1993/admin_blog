package models

import (
	"api-admin/helpers"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Link *gorm.DB
	err  error
)

// 链接数据库
func init() {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		helpers.DatabaseConfig.User,
		helpers.DatabaseConfig.Password,
		helpers.DatabaseConfig.Host,
		helpers.DatabaseConfig.Port,
		helpers.DatabaseConfig.Dbname,
		helpers.DatabaseConfig.Charset,
	)

	Link, err = gorm.Open(helpers.DatabaseConfig.Type, connectString)

	if err != nil {
		log.Fatal("Fail to connect db error : %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return helpers.DatabaseConfig.TablePrefix + defaultTableName
	}

	Link.SingularTable(true)

	Link.LogMode(helpers.DatabaseConfig.Debug)
}

type Model struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT;column:id;";json:"id"`
}

// 用户表
type Admin struct {
	UID         uint   `gorm:"primary_key;AUTO_INCREMENT;column:uid;";json:"id"`
	Username    string `gorm:"type:varchar(65);not null;unique;";json:"username"`
	Password    string `gorm:"type:char(32);not null";json:"password"`
	UserPic     string `gorm:"type:varchar(155);not null";json:"user_pic"`
	Nickname    string `gorm:"type:varchar(45);not null";json:"nickname"`
	AdminStatus int8   `gorm:"type:tinyint(1);not null;index:idx_user_status;default:1";json:"status"`
}
