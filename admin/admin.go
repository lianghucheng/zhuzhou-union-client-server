package admin

import "zhuzhou-union-client-server/models"
import (
	"github.com/qor/admin"
	"net/http"
	"zhuzhou-union-client-server/admin/article"
	"zhuzhou-union-client-server/admin/menu"
	"zhuzhou-union-client-server/admin/user"
	"zhuzhou-union-client-server/pkg/LocalI18n"
	"zhuzhou-union-client-server/admin/i18n"
)

func GetHandler() http.Handler {
	Admin := admin.New(&admin.AdminConfig{DB: models.DB, Auth: AdminAuth{}, I18n: LocalI18n.LocalI18n})

	user.SetAdmin(Admin)
	article.SetAdmin(Admin)
	menu.SetAdmin(Admin)
	i18n.SetAdmin(Admin)
	return Admin.NewServeMux("/admin")
}
