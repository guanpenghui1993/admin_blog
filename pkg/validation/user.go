package validation

// 登录验证用户参数
type UserLogin struct {
	Username string `form:"username" v:"required|length:6,20 #用户名不能为空|用户名须6-20字符"`
	Password string `form:"password" v:"required|password #密码不能为空|密码任意可见字符，长度在6~18之间"`
}

// 添加用户
type InsertUserData struct {
	Roleid   uint   `form:"roleid" v:"min:0 #角色ID错误"`
	Username string `form:"username" v:"required|length:6,20 #用户名不能为空|用户名须6-20字符"`
	Password string `form:"password" v:"required|password #密码不能为空|密码任意可见字符，长度在6~18之间"`
	UserPic  string `form:"pic" v:"required|url #头像不能为空|头像不是有效格式"`
	Nickname string `form:"nickname" v:"required|length:2,10 #昵称不能为空|昵称须2-10字符"`
}
