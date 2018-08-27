package mail

import (
"github.com/qor/admin"
"zhuzhou-union-client-server/models"
"github.com/qor/roles"
)

func SetAdmin(adminConfig *admin.Admin) {
	mailBox := adminConfig.AddResource(&models.MailBox{}, &admin.Config{Name: "主席信箱", PageCount: 10, Permission: roles.Allow(roles.Read, "admin")})

	mailBox.IndexAttrs("ID", "Title", "Content")

	mailBox.Meta(&admin.Meta{Name: "Title", Label: "标题"})
	mailBox.Meta(&admin.Meta{Name: "Content", Label: "反馈内容"})
}
