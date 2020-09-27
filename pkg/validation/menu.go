package validation

type InsertMenuData struct {
	Pid      uint   `form:"pid" v:"required #菜单级别不能为空"`
	Menuname string `form:"name" v:"required|length:2,10 #菜单名称不能为空|菜单名称6-10字符"`
	Router   string `form:"router" v:"required #菜单路由不能为空"`
	Icon     string `form:"icon" v:"required #菜单图标不能为空"`
	Sort     int    `form:"sort" v:"between:0,255 #排序值0-255"`
}

// 返回菜单结构
type MenuList struct {
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Router string `json:"router"`
	Pid    uint   `json:"pid"`
	Id     uint   `json:"id"`
	Child  []MenuList
}
