package controllers

import (
	"zhuzhou-union-client-server/utils"
	"github.com/astaxie/beego"
)

type DateControllor struct {
	Common
}

//@router /date [get]
func (this *DateControllor) Date() {
	date := utils.GetDate()
	this.ReturnSuccess("date", date)
}

//@router /weather [get]
func (this *DateControllor) Weather() {
	weather := utils.GetWeather()
	this.ReturnSuccess("weather", weather)
}

//@router /test [*]
func (this *DateControllor) Test() {
	content := this.GetString("content")
	beego.Debug(content)
	this.TplName = "test.html"
}
