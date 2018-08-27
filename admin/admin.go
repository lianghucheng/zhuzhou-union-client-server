package admin

import "zhuzhou-union-client-server/models"
import (
	"github.com/qor/admin"
	"net/http"
	"zhuzhou-union-client-server/admin/article"
	"zhuzhou-union-client-server/admin/menu"
	"zhuzhou-union-client-server/admin/user"
	"zhuzhou-union-client-server/pkg/LocalI18n"
	"zhuzhou-union-client-server/admin/category"
	"zhuzhou-union-client-server/admin/home"
	"zhuzhou-union-client-server/admin/mail"
	"zhuzhou-union-client-server/admin/staffShow"
	"github.com/qor/roles"
)

func GetHandler() http.Handler {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser.(*models.User) != nil && currentUser.(*models.User).Prioty == 1
	})

	Admin := admin.New(&admin.AdminConfig{DB: models.DB, Auth: AdminAuth{}, I18n: LocalI18n.LocalI18n})

	user.SetAdmin(Admin)
	article.SetAdmin(Admin)
	staffShow.SetAdmin(Admin)
	menu.SetAdmin(Admin)
	category.SetAdmin(Admin)
	home.SetAdmin(Admin)
	mail.SetAdmin(Admin)

	return Admin.NewServeMux("/admin")
}
