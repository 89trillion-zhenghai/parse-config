package util

import (
	"log"

	"gopkg.in/ini.v1"
)

//ReadIni 读取app.ini 配置文件，获取端口号
func ReadIni() (port string) {
	file, err := ini.Load("../conf/app.ini")
	if err != nil {
		log.Fatal(err.Error())
	}
	port = file.Section("server").Key("HttpPort").String()
	return port
}
