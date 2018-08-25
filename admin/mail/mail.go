package mail

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/roles"
)

func SetAdmin(adminConfig *admin.Admin) {
	mailBox := adminConfig.AddResource(&models.MailBox{}, &admin.Config{Name: "主席信箱", PageCount: 10,Permission:roles.Allow(roles.Read,roles.Anyone)})

	mailBox.IndexAttrs("ID", "Title", "User")
	mailBox.EditAttrs("Author", "Content", "Contact")
	mailBox.OverrideEditAttrs(func() {

	})

	mailBox.Meta(&admin.Meta{Name: "Title", Label: "标题"})
	mailBox.Meta(&admin.Meta{Name: "Author", Label: "作者"})
	mailBox.Meta(&admin.Meta{Name: "Content", Label: "反馈内容"})
	mailBox.Meta(&admin.Meta{Name: "Contact", Label: "联系方式"})
}
