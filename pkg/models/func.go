package models

import "watt/pkg/utils"

// 自动迁移
func AutoMigrateTable() {

	Link.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Role{}, &Menu{}, &Access{})

	utils.Strace("----------------------------初始化表结构结束---------------------------")
}
