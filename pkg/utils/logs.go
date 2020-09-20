package utils

import (
	"log"
	"os"
)

const (
	ERROR_PREFIX  = "[程序错误] - "
	STRACE_PREFIX = "[程序日志] - "
)

// 打印日志并终止程序
func Exit(message interface{}) {

	log.SetPrefix(ERROR_PREFIX)

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	if Setting.Common.Debug {
		log.Fatal("error -- ", message)
	}

	logFile, err := os.OpenFile(Setting.Log.ErrorLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.SetOutput(logFile)

	if err != nil {
		log.Fatal("Fail to read file -- ", err)
	}

	log.Fatal("Fail to read file -- ", message)
}

// 打印日志不终止程序
func Strace(message interface{}) {

	log.SetPrefix(STRACE_PREFIX)

	if Setting.Common.Debug {
		log.Printf("%v \n", message)
		return
	}

	logFile, err := os.OpenFile(Setting.Log.StraceLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.SetOutput(logFile)

	if err != nil {
		Exit(err)
	}

	log.Printf("%v \n", message)
}

// 记录错误日志不终止程序
func Error(message interface{}) {

	log.SetPrefix(STRACE_PREFIX)

	if Setting.Common.Debug {
		log.Printf("%v \n", message)
		return
	}

	logFile, err := os.OpenFile(Setting.Log.ErrorLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.SetOutput(logFile)

	if err != nil {
		Exit(err)
	}

	log.Printf("%v \n", message)
}
