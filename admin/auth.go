package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/session/manager"
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
)

type AdminAuth struct {
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

//从session中获得当前用户。
func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	adminUserName := manager.SessionManager.Get(c.Request, beego.AppConfig.String("adminsessionKey"))

	var user models.User

	if err := models.DB.
		Where("username=?", adminUserName).
		First(&user).
		Error; err != nil {
		return nil
	}
	return &user
}
