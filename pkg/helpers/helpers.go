package helpers

import (
	"log"
	"os"
	"time"
	"watt/pkg/config"
)

type helpers struct {
}

const (
	ERROR_PREFIX  = "[程序错误] - "
	STRACE_PREFIX = "[程序日志] - "
)

var (
	H          = newHelpers()
	strace_log = config.Conf.Log.StraceLog
	error_log  = config.Conf.Log.ErrorLog
	watt_debug = config.Conf.Common.Debug
)

func newHelpers() *helpers {
	return new(helpers)
}

// 打印日志并终止程序
func (h *helpers) Exit(message interface{}) {

	log.SetPrefix(ERROR_PREFIX)

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	if watt_debug {
		log.Fatal("error -- ", message)
	}

	logFile, err := os.OpenFile(error_log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.SetOutput(logFile)

	if err != nil {
		log.Fatal("Fail to read file -- ", err)
	}

	log.Fatal("Fail to read file -- ", message)
}

// 打印日志不终止程序
func (h *helpers) Strace(message interface{}) {

	log.SetPrefix(STRACE_PREFIX)

	if watt_debug {
		log.Printf("%v \n", message)
		return
	}

	logFile, err := os.OpenFile(strace_log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.SetOutput(logFile)

	if err != nil {
		h.Exit(err)
	}

	log.Printf("%v \n", message)
}

// 返回当前时间格式化
func (h *helpers) Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
