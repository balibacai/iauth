package main

import (
	_ "auth/routers"

	"auth/bootstrap"
	"github.com/astaxie/beego"
)

func init() {
	bootstrap.AppConfig()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
