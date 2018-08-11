package utils

import (
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"github.com/astaxie/beego"
)

// 发送短信
func SendMsg(username string){
	clnt := ypclnt.New("037ced20fb015ee88829cd3e6248aa6f")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = username
	code:=GetCode(username)
	param[ypclnt.TEXT] = "【易正网络】您的验证码是"+code
	param[ypclnt.SIGN]="【易正网络】"
	r := clnt.Sms().SingleSend(param)
	beego.Debug(r)
}
