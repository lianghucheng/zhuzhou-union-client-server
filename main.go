package main

import (
	"github.com/astaxie/beego"
	"os"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/models"
	_ "zhuzhou-union-client-server/routers"
	"github.com/qor/media/oss"
	"github.com/qor/oss/qiniu"
)

func init() {

	initArgs()
	models.Connect()
}

func main() {
	oss.Storage = qiniu.New(&qiniu.Config{
		AccessID:  beego.AppConfig.String("accessKey"),
		AccessKey: beego.AppConfig.String("secretKey"),
		Bucket:    beego.AppConfig.String("qiniuBucket"),
		Region:    beego.AppConfig.String("qiniuRegion"),
		Endpoint:  beego.AppConfig.String("qiniuUrl"),
	})
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
