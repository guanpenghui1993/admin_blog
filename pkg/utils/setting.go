package utils

import (
	"time"
)

type configYaml struct {
	Common struct {
		Debug bool `yaml:"debug,flow"`
	}

	Log struct {
		ErrorLog  string `yaml:"error_log,flow"`
		StraceLog string `yaml:"strace_log,flow"`
	}

	Jwt struct {
		Expired int    `yaml:"expired,flow"`
		Secret  string `yaml:"secret,flow"`
		Header  string `yaml:"header,flow"`
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
			Host     string `yaml:"host,flow"`
			Password string `yaml:"password,flow"`
			Port     int    `yaml:"port,flow"`
			Database int    `yaml:"database,flow"`
			Timeout  int    `yaml:"timeout,flow"`
		}
	}
}
