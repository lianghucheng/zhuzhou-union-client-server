package user

import (
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
)

type UserController struct {
	controllers.BaseController
}

//@router /userData [get]
func (this *UserController)UserData(){
	userinfo:=this.GetSession("userinfo").(*models.User)
	this.ReturnSuccess("userinfo",userinfo)
}

//@router /article/list [get]
func (this *UserController)ArticleList(){

}