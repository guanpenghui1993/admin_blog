package validation

type InsertMenuData struct {
	Pid      int    `form:"pid" v:"required #菜单级别不能为空"`
	Menuname string `form:"name" v:"required|length:6,10 #菜单名称不能为空|菜单名称6-10字符"`
	Router   string `form:"router" v:"required #菜单路由不能为空"`
	Icon     string `form:"icon" v:"required #菜单图标不能为空"`
}
