package global

import (
	"blog_server/conf"

	"gorm.io/gorm"
)

// Conf 解析yml配置信息
var (
	Conf    *conf.Config
	MysqlDB *gorm.DB
)
