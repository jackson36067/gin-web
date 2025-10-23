package models

type User struct {
	Model
	Username       string `json:"username" gorm:"size:32;unique;not null"`
	Nickname       string `json:"nickname" gorm:"size:32;not null"`
	Avatar         string `json:"avatar" gorm:"size:255;"`
	Abstract       string `json:"abstract" gorm:"size:255;"` // 简介
	RegisterSource int8   `json:"register_source"`           // 注册来源
	Password       string `json:"password" gorm:"size:120;"`
	CodeAge        int    `json:"code_age" gorm:"size:255;"` // 码龄
	Email          string `json:"email" gorm:"size:255;"`
	OpenId         string `json:"open_id" gorm:"size:64;"` // 第三方登录的唯一ip
}
