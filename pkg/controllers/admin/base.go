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

	err := c.ShouldBind(obj)

	if err != nil {
		return errors.New(err.Error())
	}

	if errs := gvalid.CheckStruct(obj, nil); errs != nil {
		return errors.New(errs.FirstString())
	}

	return nil
}

// 根据token获取用户ID
func (b *BaseController) getuid(c *gin.Context) int {

	uid, _ := utils.Parse(c.GetHeader(utils.Setting.Jwt.Header))

	return uid
}
