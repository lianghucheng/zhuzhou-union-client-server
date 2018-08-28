package mail

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/roles"
)

func SetAdmin(adminConfig *admin.Admin) {
	mailBox := adminConfig.AddResource(&models.MailBox{}, &admin.Config{Name: "主席信箱", PageCount: 10, Permission: roles.Allow(roles.Read, roles.Anyone)})

	mailBox.IndexAttrs("ID", "Title", "CreatedAt", "Ip")

	mailBox.Meta(&admin.Meta{Name: "Title", Label: "标题"})
	mailBox.Meta(&admin.Meta{Name: "CreatedAt", Label: "发表时间"})
	mailBox.Meta(&admin.Meta{Name: "Ip", Label: "Ip地址"})

	mailBox.EditAttrs("ID", "Title", "Content")

	mailBox.Meta(&admin.Meta{Name: "Content", Type: "kindeditor", Label: "内容"})
	mailBox.SearchAttrs("Title", "Content")
}
