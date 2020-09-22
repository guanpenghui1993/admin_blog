package models

import (
	"fmt"
	"watt/pkg/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Link *gorm.DB
	err  error
)

// 链接数据库
func init() {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%ds",
		utils.Setting.Database.Mysql.User,
		utils.Setting.Database.Mysql.Password,
		utils.Setting.Database.Mysql.Host,
		utils.Setting.Database.Mysql.Port,
		utils.Setting.Database.Mysql.Dbname,
		utils.Setting.Database.Mysql.Charset,
		utils.Setting.Database.Mysql.Timeout,
	)

	Link, err = gorm.Open(utils.Setting.Database.Mysql.Dbtype, connectString)

	if err != nil {
		utils.Exit(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return utils.Setting.Database.Mysql.Prefix + defaultTableName
	}

	Link.SingularTable(true)

	Link.LogMode(utils.Setting.Common.Debug)

	Link.DB().SetMaxIdleConns(10)

	Link.DB().SetMaxOpenConns(100)
}
