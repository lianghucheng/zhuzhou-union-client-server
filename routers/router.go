package routers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/controllers/article"
	"zhuzhou-union-client-server/controllers/auth"
	"zhuzhou-union-client-server/controllers/home"
	"zhuzhou-union-client-server/controllers/user"
	"github.com/dchest/captcha"
)

func init() {
	beego.Include(
		&admin.FileUploadController{},
		&admin.LoginController{},
		&controllers.DateControllor{},
		&auth.Controller{},
		&home.Controller{},
		&user.Controller{},
		&article.Controller{},
		&controllers.CommonController{},
	)
	beego.Router("/api/ueditor_controller", &controllers.Ueditor{}, "*:U_Controller")
	beego.Handler("/api/image/captcha/*.png", captcha.Server(90, 40))
	beego.SetStaticPath("/image/kindeditor/upload", "/upload")
	beego.SetStaticPath("/UploadFiles", "UploadFiles")
}
