package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/controllers/article"
	"zhuzhou-union-client-server/controllers/auth"
	"zhuzhou-union-client-server/controllers/home"
	"zhuzhou-union-client-server/controllers/user"
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
	beego.Router("/api/ueditor_controller", &controllers.Ueditor{}, "*:U_Controller")
	beego.SetStaticPath("/admin/assets/javascripts/ueditor/*", "app/views/qor/assets/javascripts/ueditor/")
	beego.SetStaticPath("/UploadFiles", "UploadFiles")

}
