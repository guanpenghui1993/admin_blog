package models

type demo struct {
	Nickname string `json:"nickname"`
	Uid      uint   `json:"uid"`
}

func GetUserList() []demo {

	var demo []demo

	Link.Table("blog_admin").Select([]string{"nickname", "uid"}).Scan(&demo)
	// var admin []Admin

	// err := Link.Find(&admin).Error

	// if err != nil {
	// log.Fatnl(err)
	// }

	return demo
}
