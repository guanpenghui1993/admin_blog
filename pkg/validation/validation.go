package validation

type BaseValid struct {
	Page int `form:"page" v:"required|min:1|max:100 #分页码为空|分页码错误|分页最大不超100"`
	Size int `form:"size" v:"required|min:10|max:50 #数据大小为空|数据最小10|数据最大50"`
}

type BaseID struct {
	ID int `form:"id" v:"required|min:1 #id不能为空|id错误"`
}

type BaseIdStatus struct {
	ID     int `form:"id" v:"required|integer #id不能为空|id错误"`
	Status int `form:"status" v:"required|in:0,1,-1 #状态不能为空|状态错误"`
}

type InsertMenuData struct {
	Pid      uint   `form:"pid" v:"required #菜单级别不能为空"`
	Menuname string `form:"name" v:"required|length:2,10 #菜单名称不能为空|菜单名称2-10字符"`
	Router   string `form:"router" v:"required #菜单路由不能为空"`
	Icon     string `form:"icon" v:"required #菜单图标不能为空"`
	Sort     int    `form:"sort" v:"between:0,255 #排序值0-255"`
}

// 返回菜单结构
type MenuList struct {
	Name   string     `json:"name"`
	Icon   string     `json:"icon"`
	Router string     `json:"router"`
	Pid    uint       `json:"pid"`
	Id     uint       `json:"id"`
	Child  []MenuList `json:"child"`
}

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

// 角色
type RoleData struct {
	Rolename string `form:"name" v:"required|length:2,10 #角色名称不能为空|角色名称2-10字符"`
	Roledesc string `form:"desc" v:"required|length:5,45 #角色描述不能为空|角色描述5-45字符"`
}
