package i18n

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/pkg/LocalI18n"
)

func SetAdmin(adminConfig *admin.Admin) {
	Locali18n := adminConfig.AddResource(LocalI18n.LocalI18n, &admin.Config{Name: "国际化", PageCount: 10})

	Locali18n.SearchAttrs("Translation")
}
