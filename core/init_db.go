package core

import (
	"blog_server/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func InitMysql() (db *gorm.DB) {
	dsn := global.Conf.DB.Mysql
	dsn1 := global.Conf.DB.Mysql1
	db, err := gorm.Open(mysql.Open(dsn.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logrus.Fatalf("连接数据库失败: %s", err)
	}
	// 设置连接池大小
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功")
	// 配置读写数据库(分库)
	if !dsn1.Empty() {
		err = db.Use(dbresolver.Register(dbresolver.Config{
			// use `db2` as sources, `db3`, `db4` as replicas
			Sources:  []gorm.Dialector{mysql.Open(dsn1.DSN())}, // 写
			Replicas: []gorm.Dialector{mysql.Open(dsn.DSN())},  // 读
			// sources/replicas load balancing policy
			Policy: dbresolver.RandomPolicy{},
			// print sources/replicas mode in logger
			TraceResolverMode: true,
		}))
		if err != nil {
			logrus.Fatalf("读写数据库配置出错: %s", err)
		}
		logrus.Infof("读写数据库配置成功")
	}
	return db
}
