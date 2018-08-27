package controllers

import "zhuzhou-union-client-server/models"

type MailController struct {
	Common
}

//@router /mail     [*]
func (this *MailController) Index() {
	this.TplName = "mail.html"
}

//@router /mail/add [post]
func (this *MailController) Add() {
	title:=this.GetString("title")
	content:=this.GetString("content")

	var mail models.MailBox

	mail.Title=title
	mail.Content=content

	if err:=models.DB.Save(&mail).Error;err!=nil{
		this.Abort("500")
		return
	}
	this.ReturnSuccess()
}


