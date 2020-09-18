package config

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	Conf configYaml
	file string
)

func init() {

	flag.StringVar(&file, "f", "config/application.yaml", "应用系统配置文件")

	flag.Parse()

	config, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("Fail to read file: %v - ", err)
	}

	yaml.Unmarshal(config, &Conf)
}
