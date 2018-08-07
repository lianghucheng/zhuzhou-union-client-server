package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:FileUploadController"] = append(beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:FileUploadController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/image/kindeditor/upload`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:LoginController"] = append(beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:LoginController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/auth/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:LoginController"] = append(beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:LoginController"],
		beego.ControllerComments{
			Method: "LoginSubmit",
			Router: `/auth/login/submit`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:LoginController"] = append(beego.GlobalControllerRouter["zhuzhou-union-client-server/admin:LoginController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/auth/logout`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
