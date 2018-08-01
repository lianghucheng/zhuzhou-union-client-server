package main

import (
	"github.com/astaxie/beego"
	"github.com/qor/admin"
	"os"
	"zhuzhou-union-client-server/models"
	_ "zhuzhou-union-client-server/routers"
)

func init() {
	initArgs()
	models.Connect()
}

func main() {
	Admin := admin.New(&admin.AdminConfig{DB: models.DB})

	Admin.AddResource(&models.User{})
	Admin.AddResource(&models.Category{})

	article := Admin.AddResource(&models.Article{})
	article.Meta(&admin.Meta{Name: "Content", Type: "rich_editor"})
	article.Meta(&admin.Meta{Name: "Status", Config: &admin.SelectOneConfig{Collection: []string{"显示", "不显示", "审核中"}}})

	beego.Handler("/admin", Admin.NewServeMux("/admin"), true)
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
