package conf

import "fmt"

type DB struct {
	Mysql  Mysql `yaml:"mysql"`
	Mysql1 Mysql `yaml:"mysql1"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Debug    bool   `yaml:"debug"` // 打印全部的日志
}

func (m Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2FShanghai", m.Username, m.Password, m.Host, m.Port, m.Dbname)
}

func (m Mysql) Empty() bool {
	return m.Username == "" && m.Password == "" && m.Host == "" && m.Port == ""
}
