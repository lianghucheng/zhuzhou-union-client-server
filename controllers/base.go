package controllers

import (
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
		this.Userinfo = u
		this.Data["user"] = u
	} else {
		this.Ctx.Redirect(302, "/")
		return
	}
}
