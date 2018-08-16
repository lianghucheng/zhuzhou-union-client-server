package controllers

import (
	"zhuzhou-union-client-server/utils"
)

type DateControllor struct {
	Common
}

//@router /date [get]
func (this *DateControllor) Date() {
	date:=utils.GetDate()
	this.ReturnSuccess("date",date)
}

//@router /weather [get]
func (this *DateControllor)Weather(){
	weather:=utils.GetWeather()
	this.ReturnSuccess("weather",weather)
}
