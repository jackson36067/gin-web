package core

import (
	"blog_server/conf"
	"blog_server/flags"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadConf 读取配置文件
func ReadConf() (c *conf.Config) {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	c = &conf.Config{}
	err = yaml.Unmarshal(byteData, &c)
	if err != nil {
		panic(fmt.Sprintf("yml配置文件格式错误 %s", err))
	}
	fmt.Printf("读取配置文件%s成功\n", flags.FlagOptions.File)
	return
}
