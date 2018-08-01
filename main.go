package main

import (
	"github.com/astaxie/beego"
	"os"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/models"
	_ "zhuzhou-union-client-server/routers"
)

func init() {
	initArgs()
	models.Connect()
}

func main() {
	beego.Handler("/admin", admin.GetHandler(), true)
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.SyncDB()
			os.Exit(0)
		}
	}
}
