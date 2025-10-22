package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var confPath = "settings.yml"

type Config struct {
	System struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"system"`
}

// ReadConf 读取配置文件
func ReadConf() {
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	var conf Config
	err = yaml.Unmarshal(byteData, &conf)
	if err != nil {
		panic(fmt.Sprintf("yml配置文件格式错误 %s", err))
	}
	fmt.Println(conf)
}
