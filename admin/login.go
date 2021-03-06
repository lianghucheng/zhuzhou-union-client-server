package admin

import (
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
	"zhuzhou-union-client-server/controllers"
	"github.com/qor/session/manager"
	"github.com/astaxie/beego"
)

type LoginController struct {
	controllers.Common
}

//@router /auth/login [*]
func (this *LoginController) Index() {
	this.TplName = "admin/login.html"
}

//@router /auth/login/submit [*]
func (this *LoginController) LoginSubmit() {
	username := this.GetString("username")
	password := this.GetString("password")

	var user models.User

	enPassword := utils.Md5(password)
	if err := models.DB.
		Where("username=? and password=?", username, enPassword).
		First(&user).
		Error;
		err != nil {
		this.ReturnJson(10001, "用户名或密码错误")
		return
	}
	if user.Prioty == 3 {
		beego.Debug("该用户不是管理员")
		this.ReturnJson(10003,"您不是管理员")
	}
	manager.SessionManager.Add(this.Ctx.ResponseWriter, this.Ctx.Request, beego.AppConfig.String("adminsessionKey"), user.Username)
	this.ReturnSuccess()

}

//@router /auth/logout [*]
func (this *LoginController) Logout() {

	manager.SessionManager.Pop(this.Ctx.ResponseWriter, this.Ctx.Request, beego.AppConfig.String("adminsessionKey"))
	//设置返回对象。
	this.Ctx.Redirect(302, "/auth/login")
	return
}
