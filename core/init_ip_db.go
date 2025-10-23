package core

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var IpDB *geoip2.Reader

// InitIPDB 读取ip地址数据库文件
func InitIPDB() {
	db, err := geoip2.Open("init/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatalf("ip地址数据库加载失败, %s", err)
	}
	IpDB = db
}

// GetIpAddress 根据IP地址以及IP地址数据库文件解析获取地址
func GetIpAddress(ip string) (string, error) {
	pIp := net.ParseIP(ip)
	if pIp == nil {
		return "", fmt.Errorf("无效的IP地址: %s", ip)
	}

	record, err := IpDB.City(pIp)
	if err != nil {
		return "", err
	}

	country := record.Country.Names["zh-CN"]
	province := ""
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}
	city := record.City.Names["zh-CN"]
	if city != "" {
		return fmt.Sprintf("%s %s %s", country, province, city), nil
	}
	if province != "" {
		return fmt.Sprintf("%s %s", country, province), nil
	}
	return country, nil
}
