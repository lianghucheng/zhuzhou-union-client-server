package routers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/controllers/user"
	"zhuzhou-union-client-server/controllers"
)

func init() {
	beego.Include(&admin.FileUploadController{})
	beego.Include(&admin.LoginController{})
	beego.Include(&user.UserController{})
	beego.Include(&controllers.DateControllor{}, &controllers.HomeController{})
	beego.SetStaticPath("/image/kindeditor/upload", "/upload")

}
