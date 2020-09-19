package utils

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	Setting configYaml
	file    string
)

func init() {

	flag.StringVar(&file, "f", "config/application.yaml", "应用系统配置文件")

	flag.Parse()

	config, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("Fail to read file: %v - ", err)
	}

	yaml.Unmarshal(config, &Setting)
}
