package main

import (
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"
)

func main() {
	// 解析命令行参数, 判断加载哪个配置文件
	flags.Parse()
	// 读取配置文件
	global.Conf = core.ReadConf()
	// 配置Logrus
	core.InitLogrus()
	// 连接数据库
	global.MysqlDB = core.InitMysql()
	// 数据库迁移
	flags.Run()
}
