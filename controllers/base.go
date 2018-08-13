package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
)

type BaseController struct {
	Common
}

func (this *BaseController) Prepare() {
	if u, ok := this.GetSession("userinfo").(*models.User); ok {

		if err := models.DB.Where("id = ?", u.ID).First(u).Error; err != nil {
			this.Ctx.Redirect(302, "/")
			return
		}
		this.Userinfo=u
	} else {
		this.Ctx.Redirect(302, "/")
		return
	}
}


