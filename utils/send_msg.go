package utils

import (
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"github.com/astaxie/beego"
)

// 发送短信
func SendMsg(username string){
	clnt := ypclnt.New(beego.AppConfig.String("yun_pian_apikey"))
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = username
	code:=GetCode(username)
	param[ypclnt.TEXT] = beego.AppConfig.String("msg_template")+code
	r := clnt.Sms().SingleSend(param)
	beego.Debug(r)
}
