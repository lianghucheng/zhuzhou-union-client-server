package admin

import (
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
	"zhuzhou-union-client-server/controllers"
)

type LoginController struct {
	controllers.Common
}

//@router /auth/login [*]
func (this *LoginController) Index() {
	this.TplName = "index.html"
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
	this.SetSession("adminuser", user)
	this.ReturnSuccess()

}
