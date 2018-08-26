package staffShow

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/roles"
)

func SetAdmin(adminConfig *admin.Admin) {
	staff := adminConfig.AddResource(&models.StaffShow{}, &admin.Config{Name: "职工show管理", PageCount: 10,Permission:roles.Allow(roles.Read,roles.Anyone).Allow(roles.Update,roles.Anyone)})

	//

	staff.SearchAttrs("Name", "ID", "URL", "Category", "Higher", "Sequence")

	staff.IndexAttrs("ID", "Title", "User", "Category")
	staff.EditAttrs("ID", "Title", "User", "Category")
	staff.NewAttrs("ID", "Title", "User", "Category")

	staff.Meta(&admin.Meta{Name: "Title", Label: "标题"})
	staff.Meta(&admin.Meta{Name: "User", Label: "用户", Config: &admin.SelectOneConfig{Placeholder: "选择选项"}})
	staff.Meta(&admin.Meta{Name: "Category", Label: "类别"})

}
