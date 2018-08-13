package utils

import (
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"github.com/astaxie/beego"
)

// 发送短信
func SendMsg(username string, code string) {
	clnt := ypclnt.New(beego.AppConfig.String("yun_pian_apikey"))
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = username
	param[ypclnt.SIGN] = "【易正网络】"
	param[ypclnt.TEXT] = beego.AppConfig.String("msg_template") + code
	r := clnt.Sms().SingleSend(param)
	beego.Debug(r)
}
