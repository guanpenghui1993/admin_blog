package validation

type BaseValid struct {
	Page int `form:"page" v:"required|min:1 #分页码为空|分页码错误"`
	Size int `form:"size" v:"required|min:10|max:50 #数据大小为空|数据最小10|数据最大50"`
}
