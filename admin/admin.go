package admin

import "zhuzhou-union-client-server/models"
import (
	"github.com/astaxie/beego"
	"github.com/qor/admin"
)

func init() {
	Admin := admin.New(&admin.AdminConfig{DB: models.DB})

	Admin.AddResource(&models.User{})
	Admin.AddResource(&models.Category{})
	article := Admin.AddResource(&models.Article{})

	article.Meta(&admin.Meta{Name: "Content", Type: "rich_editor"})
	article.Meta(&admin.Meta{Name: "Status", Config: &admin.SelectOneConfig{Collection: []string{"显示", "不显示", "审核中"}}})

	beego.Handler("/admin", Admin.NewServeMux("/admin"), true)
}
