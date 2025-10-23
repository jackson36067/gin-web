package flags

import (
	"blog_server/global"
	"blog_server/models"

	"github.com/sirupsen/logrus"
)

func MigrateDB() {
	err := global.MysqlDB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		logrus.Fatalf("数据库迁移失败 %s", err)
		return
	}
	logrus.Infof("数据库迁移成功")
}
