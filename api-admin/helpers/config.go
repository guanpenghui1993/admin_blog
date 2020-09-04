package helpers

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

const configPath string = "config/application.ini"

type Path struct {
	Logs string
}

type Server struct {
	Port         uint
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type     string
	User     string
	Password string
	Dbname   string
	Host        string
	Port        uint
	Charset     string
	TablePrefix string
}

var (
	PathConfig     = new(Path)
	ServerConfig   = new(Server)
	DatabaseConfig = new(Database)
	cfg            *ini.File
	err            error
)

func init() {

	cfg, err = ini.LoadSources(ini.LoadOptions{AllowBooleanKeys: true}, configPath)

	if err != nil {
		log.Fatal("Fail to read file: %v -- ", err)
	}

	mapTo("path", PathConfig)
	mapTo("server", ServerConfig)
	mapTo("database", DatabaseConfig)
}

func mapTo(section string, paramNamePointer interface{}) {

	err := cfg.Section(section).MapTo(paramNamePointer)

	if err != nil {
		log.Fatalf("Config MapTo Setting err: %v", err)
	}
}
