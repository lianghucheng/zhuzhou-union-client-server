package controllers

import (
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego"
)

type MailController struct {
	Common
}

func (this *MailController) Prepare() {
	this.initMenu()
	this.initFooter()

	var code models.QrCode
	if err := models.DB.First(&code).Error; err != nil {
		beego.Error("没有发现二维码")
	}

	this.Data["Code"] = code


	if u, ok := this.GetSession("userinfo").(*models.User); ok {

		if err := models.DB.Where("id = ?", u.ID).First(u).Error; err != nil {
			this.Ctx.Redirect(302, "/user/login")
			return
		}
		this.Userinfo = u
		this.Data["User"] = u
	} else {
		this.Ctx.Redirect(302, "/user/login")
		return
	}
}

//@router /mail     [*]
func (this *MailController) Index() {
	this.TplName = "mail.html"
}

//@router /mail/add [post]
func (this *MailController) Add() {
	title := this.GetString("title")
	content := this.GetString("content")

	var mail models.MailBox

	mail.Title = title
	mail.Content = content
	mail.Ip = this.Ctx.Input.IP()
	if err := models.DB.Save(&mail).Error; err != nil {
		this.Abort("500")
		return
	}
	this.ReturnSuccess()
}
