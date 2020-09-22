package validation

// 登录验证用户参数
type UserLogin struct {
	Username string `form:"username" v:"required|length:6,20 #用户名不能为空|用户名须6-20字符"`
	Password string `form:"password" v:"required|length:6,20 #密码不能为空|密码须6-20字符"`
}
