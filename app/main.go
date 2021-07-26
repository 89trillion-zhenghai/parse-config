package main

import (
	"parse-config/app/http"
	"parse-config/util"
)

func main() {
	util.Init()
	port := util.ReadIni()
	http.InitServer(port)
}
