package models

type Admin struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;column:uid";json:"id"`
	Username    string `gorm:"type:varchar(65);not null";json:"username"`
	Password    string `gorm:"type:char(32);not null";json:"password"`
	UserPic     string `gorm:"type:varchar(155);not null";json:"user_pic"`
	Nickname    string `gorm:"type:varchar(45);not null";json:"nickname"`
	AdminStatus int8   `gorm:"type:tinyint(1);not null;default:1";json:"status"`
}

func (a *Admin) Select() (interface{}, error) {

	type Result struct {
		Nickname string `json:"nickname"`
		Uid      int
		Username string
	}

	var adminList []Result

	link.Table("blog_admin").Select([]string{"uid", "nickname", "username"}).Scan(&adminList)

	// err := link.Where("admin_status = ?", 1).Find(&adminList).Error

	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return nil, err
	// }

	return adminList, nil
}
