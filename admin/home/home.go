package home

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
)

func SetAdmin(adminConfig *admin.Admin) {
	home := adminConfig.AddResource(&models.Home{}, &admin.Config{Name: "首页管理"})
	//对增删查改的局部显示
	home.IndexAttrs("ID", "Category", "IndexArticle", "Position")
	home.EditAttrs("Category", "IndexArticle", "Position")
	home.NewAttrs("Category", "IndexArticle", "Position")

	home.Meta(&admin.Meta{Name: "Category", Label: "首页分类"})
	home.Meta(&admin.Meta{Name: "Position", Label: "具体位置"})
	home.Meta(&admin.Meta{Name: "Layout", Label: "模块位置"})
	home.Meta(&admin.Meta{Name: "IndexArticle", Label: "单个分类置顶文章"})

}
