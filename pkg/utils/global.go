package utils

import "time"

// 全局配置
var Setting ConfigYaml

// 全局常量
const (
	SUCCESS        = 200
	ERROR          = 201
	HEADER_ERROR   = 202
	PARAM_ERROR    = 204
	FORIDDEN_ERROR = 403
	SERVER_ERROR   = 500
	SUPER_ROLE     = 1
	ERROR_PREFIX   = "[程序错误] - "
	STRACE_PREFIX  = "[程序日志] - "
)

// 全局响应返回格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 全局配置结构
type ConfigYaml struct {
	Common struct {
		Debug bool `yaml:"debug,flow"`
	}

	Log struct {
		ErrorLog  string `yaml:"error_log,flow"`
		StraceLog string `yaml:"strace_log,flow"`
	}

	Jwt struct {
		Expired time.Duration `yaml:"expired,flow"`
		Secret  string        `yaml:"secret,flow"`
		Issuer  string        `yaml:"issuer,flow"`
		Header  string        `yaml:"header,flow"`
	}

	Cross struct {
		Domain string `yaml:"domain,flow"`
		Header string `yaml:"header,flow"`
		Method string `yaml:"method,flow"`
	}

	Http struct {
		Port         int           `yaml:"port,flow"`
		ReadTimeout  time.Duration `yaml:"readTimeout,flow"`
		WriteTimeout time.Duration `yaml:"writeTimeout,flow"`
	}

	Database struct {
		Mysql struct {
			Dbtype   string `yaml:"type,flow"`
			Host     string `yaml:"host,flow"`
			Port     int    `yaml:"port,flow"`
			User     string `yaml:"user,flow"`
			Password string `yaml:"password,flow"`
			Dbname   string `yaml:"dbname,flow"`
			Charset  string `yaml:"charset,flow"`
			Timeout  int    `yaml:"timeout,flow"`
			Prefix   string `yaml:"prefix,flow"`
		}
		Redis struct {
			Host     string        `yaml:"host,flow"`
			Password string        `yaml:"password,flow"`
			Port     int           `yaml:"port,flow"`
			Database int           `yaml:"database,flow"`
			Timeout  time.Duration `yaml:"timeout,flow"`
		}
	}
}
