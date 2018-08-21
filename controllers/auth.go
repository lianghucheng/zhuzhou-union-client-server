package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
)

type AuthController struct {
	Common
}

//@router /user/pwd_login [*]
func (this *AuthController) PwdLogin() {
	this.TplName = "user/pwdLogin.html"
}

//@router /user/sms_login [*]
func (this *AuthController) SmsLogin(){
	this.TplName = "user/smsLogin.html"
}

//@router /user/register [*]
func (this *AuthController) Register() {
	this.TplName = "user/register.html"
}

//@router /api/auth/login [*]
func (this *AuthController) LoginSubmit() {
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "" {
		this.ReturnJson(10001, "用户名不能为空")
		return
	}

	if !MobileRegexp(username) {
		this.ReturnJson(10002, "请输入正确的手机号码")
	}
	u := models.User{}
	if models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10003, "用户名不存在")
	}

	if password == "" {
		this.ReturnJson(10004, "您的输入不完整")
	} else {
		//this.VerityCode(username)
		if models.DB.Where("username = ? and password = ?", username, utils.Md5(password)).Find(&u).RecordNotFound() {
			this.ReturnJson(10005, "用户名或者密码错误")
		} else {
			this.SetSession("userinfo", &u)
			this.ReturnSuccess()
		}
	}
}

//@router /api/auth/register [*]
func (this *AuthController) RegisterSubmit() {
	username := this.GetString("username")
	password := this.GetString("password")
	if !MobileRegexp(username) {
		this.ReturnJson(10001, "请输入正确的手机号码")
	}
	if password == "" {
		this.ReturnJson(10002, "密码不能为空")
	}

	//验证码验证
	//this.VerityCode(username)

	u := models.User{}
	if !models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10005, "该用户已经注册")
	}
	beego.Debug(username, password)
	u.Username = username
	u.Password = utils.Md5(password)
	u.Prioty = 3

	if err := models.DB.Create(&u).Error; err != nil {
		this.ReturnJson(10006, "注册用户失败")
	}

	this.SetSession("userinfo", &u)
	this.ReturnSuccess()
}

//@router /api/auth/logout [*]
func (this *AuthController) Logout() {
	this.DelSession("userinfo")
	this.ReturnSuccess()
}

//@router /api/auth/send/sms [post]
func (this *AuthController) SendSms() {
	username := this.GetString("username")
	if username == "" {
		this.ReturnJson(10001, "手机号不能为空")
		return
	}
	if !MobileRegexp(username) {
		this.ReturnJson(10001, "请输入正确的手机号码")
		return
	}
	code := string(utils.Krand(6, 0))
	this.SetSession(username, code)
	go utils.SendMsg(username, code)
	this.ReturnSuccess()
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
