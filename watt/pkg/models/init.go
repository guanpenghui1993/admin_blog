package models

import (
	"fmt"
	"watt/pkg/config"
	"watt/pkg/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Link *gorm.DB
	err  error
)

// 链接数据库
func init() {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%vs",
		config.Conf.Database.Mysql.User,
		config.Conf.Database.Mysql.Password,
		config.Conf.Database.Mysql.Host,
		config.Conf.Database.Mysql.Port,
		config.Conf.Database.Mysql.Dbname,
		config.Conf.Database.Mysql.Charset,
		config.Conf.Database.Mysql.Timeout,
	)

	Link, err = gorm.Open(config.Conf.Database.Mysql.Dbtype, connectString)

	if err != nil {
		helpers.H.Exit(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Conf.Database.Mysql.Prefix + defaultTableName
	}

	Link.SingularTable(true)

	Link.LogMode(config.Conf.Common.Debug)
}
