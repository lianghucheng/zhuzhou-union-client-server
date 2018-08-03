package admin

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

//@router /auth/login [*]
func (this *LoginController) Index() {
	fmt.Println("this is login index")
	this.TplName = "login.html"
}

//@router /auth/login/submit [*]
func (this *LoginController) LoginSubmit() {
	this.Ctx.Redirect(302, "/admin")
}
