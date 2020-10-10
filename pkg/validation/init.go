package validation

// import (
// 	"errors"

// 	"github.com/gogf/gf/util/gvalid"
// )

// 判断1,3,4是否为int类型id字符串
// gvalid.RegisterRule("stringID", func(rule string, value interface{}, message string, params map[string]interface{}) error {

// 	var idArray = strings.Split(value,",")

// 	if len(idArray) <= 0 {
// 		return errors.New("参数需以逗号相隔的字符串")
// 	}

// 	for _,val := range idArray {

// 		if valInt,_ := strconv.Atoi(val); valInt <= 0 {

// 			return errors.New(message)
// 		}
// 	}

// 	return nil
// })
