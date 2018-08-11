package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
)

type AuthController struct {
	beego.Controller
	Common
}

//@router /api/auth/login [*]
func (this *AuthController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	randNum := this.GetString("randNum")

	if username == "" {
		this.ReturnJson(10001, "用户名不能为空")
		return
	}

	if !MobileRegexp(username) {
		this.ReturnJson(10002, "请输入正确的手机号码")
		return
	}
	var u *models.User
	if models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10003, "用户名不存在")
		return
	}

	if password == "" {
		if randNum == "" {
			this.ReturnJson(10004, "您的输入不完整")
			return
		} else {
			//验证验证码是否正确
		}
	} else {
		if models.DB.Where("username = ? and password = ?", username, utils.Md5(password)).RecordNotFound() {
			this.ReturnJson(10005, "用户名或者密码错误")
			return
		}
		this.SetSession("userinfo", &u)
		this.ReturnSuccess()
	}

}

//@router /api/auth/register [*]
func (this *AuthController) Register() {
	username := this.GetString("username")
	password := this.GetString("password")
	randNum := this.GetString("randNum")
	if !MobileRegexp(username) {
		this.ReturnJson(10001, "请输入正确的手机号码")
		return
	}
	if password == "" {
		this.ReturnJson(10002, "密码不能为空")
		return
	}

	if randNum == "" {
		this.ReturnJson(10003, "验证码不能为空")
		return
	}

	//验证码验证
	if randNum != "" {
		this.ReturnJson(10004, "验证码错误")
		return
	}

	var u *models.User
	if !models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10005, "该用户已经注册")
		return
	}

	u.Username = username
	u.Password = utils.Md5(password)

	if err := models.DB.Create(&u).Error; err != nil {
		this.ReturnJson(10006, "注册用户失败")
		return
	}

	this.SetSession("userinfo", &u)
	this.ReturnSuccess()
}

//@router /api/auth/logout [*]
func (this *AuthController) Logout() {
	this.DelSession("userinfo")
	this.ReturnSuccess()
}

//@router /api/auth/send/sms [*]
func (this *AuthController) SendSms() {

}

//传入手机号码，返回处理结果的字符串类型信息
func MobileRegexp(mobile string) bool {
	matched, err := regexp.MatchString(beego.AppConfig.String("yidong"), mobile)
	if matched && err == nil {

		return true
	}
	matched, err = regexp.MatchString(beego.AppConfig.String("liantong"), mobile)
	if matched && err == nil {
		return true
	}
	matched, err = regexp.MatchString(beego.AppConfig.String("dianxin"), mobile)
	if matched && err == nil {
		return true
	}
	return false
}
