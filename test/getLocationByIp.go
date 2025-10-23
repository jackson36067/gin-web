package main

import (
	"blog_server/core"
	"fmt"
)

func main() {
	geoIp()
}

func geoIp() {
	core.InitIPDB()
	addr, err := core.GetIpAddress("58.247.0.18")
	fmt.Println(addr, err)
}
