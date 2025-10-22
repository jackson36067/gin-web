package main

import (
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"

	"github.com/sirupsen/logrus"
)

func main() {
	// 解析命令行参数, 判断加载哪个配置文件
	flags.Parse()
	// 读取配置文件
	global.Conf = core.ReadConf()
	// 配置logrus
	core.InitLogrus()
	//fmt.Println(core.Conf)
	logrus.Warnf("xxx")
	logrus.Infof("yyy")
	logrus.Errorf("zzz")
	logrus.Debug("ccc")
}
