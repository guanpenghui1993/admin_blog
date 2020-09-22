package admin

import (
	"errors"
	"watt/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

type BaseController struct {
}

// 统一返回json
func (*BaseController) json(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, utils.Response{code, message, data})
}

// 统一校验参数
func (b *BaseController) valid(c *gin.Context, obj interface{}) error {

	c.ShouldBind(obj)

	if err := gvalid.CheckStruct(obj, nil); err != nil {
		return errors.New(err.FirstString())
	}

	return nil
}

// 根据立牌获取用户ID
func (b *BaseController) getuid(c *gin.Context) int {

	uid, _ := utils.Parse(c.GetHeader(utils.Setting.Jwt.Header))

	return uid
}
