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
