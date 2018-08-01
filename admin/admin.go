package admin

import "zhuzhou-union-client-server/models"
import (
	"github.com/qor/admin"
	"net/http"
	"zhuzhou-union-client-server/admin/article"
	"zhuzhou-union-client-server/admin/menu"
	"zhuzhou-union-client-server/admin/user"
)

func GetHandler() http.Handler {
	Admin := admin.New(&admin.AdminConfig{DB: models.DB})

	user.SetAdmin(Admin)
	article.SetAdmin(Admin)
	menu.SetAdmin(Admin)

	return Admin.NewServeMux("/admin")
}
