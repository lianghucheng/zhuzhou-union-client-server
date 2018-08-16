package routers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/controllers"
)

func init() {
	beego.Include(
		&admin.FileUploadController{},
		&admin.LoginController{},
		&controllers.UserController{},
		&controllers.DateControllor{},
		&controllers.AuthController{},
		&controllers.HomeController{},
		&controllers.CommonController{},
	)
	beego.Router("/api/ueditor_controller", &controllers.Ueditor{}, "*:U_Controller")
	beego.SetStaticPath("/image/kindeditor/upload", "/upload")
}
