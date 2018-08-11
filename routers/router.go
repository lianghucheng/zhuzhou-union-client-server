package routers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/admin"
)

func init() {
	beego.Include(&admin.FileUploadController{})
	beego.Include(&admin.LoginController{})
	beego.SetStaticPath("/image/kindeditor/upload", "/upload")

}
