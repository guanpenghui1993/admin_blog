package utils

const (
	SUCCESS      = 200
	ERROR        = 201
	SERVER_ERROR = 500
	PARAM_ERROR  = 204
)

var message = map[int]string{
	SUCCESS:      "success",
	ERROR:        "error",
	SERVER_ERROR: "内部异常",
	PARAM_ERROR:  "参数错误",
}

func Getmessage(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	}
	return ""
}
