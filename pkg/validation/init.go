package validation

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gogf/gf/util/gvalid"
)

// 注册规则
func init() {
	gvalid.RegisterRule("stringID", stringTid)
}

// 判断1,3,4是否为int类型id字符串
func stringTid(value interface{}, message string, params map[string]interface{}) error {

	valstring := value.(string)

	idArray := strings.Split(valstring, ",")

	if len(idArray) <= 0 {
		return errors.New("参数需以逗号相隔的字符串")
	}

	for _, val := range idArray {

		if valInt, _ := strconv.Atoi(val); valInt <= 0 {

			return errors.New(message)
		}
	}

	return nil
}
