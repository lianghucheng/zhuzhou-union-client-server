package auth

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
	"github.com/dchest/captcha"
	"time"
)

type Controller struct {
	controllers.Common
}

// @router /api/common/captcha [*]
func (this *Controller) GetCaptcha() {
	captchaId := captcha.NewLen(4)
	result := make(map[string]interface{})
	result["status"] = 10000
	result["src"] = "/api/image/captcha/" + captchaId + ".png"
	result["id"] = captchaId
	beego.Debug(result)
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

//@router /user/login [*]
func (this *Controller) Login() {
	this.TplName = "auth/login.html"
}

//@router /user/smslogin [*]
func (this *Controller) SmsLogin() {
	this.TplName = "auth/smslogin.html"
}

//function_map 1207

//@router /user/register [*]
func (this *Controller) Register() {
	this.TplName = "auth/register.html"
}

//@router /api/user/login [*]
func (this *Controller) LoginSubmit() {
	this.CaptchaInterceptor()
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "" {
		this.ReturnJson(10001, "用户名不能为空")
		return
	}

	if !utils.MobileRegexp(username) {
		this.ReturnJson(10002, "请输入正确的手机号码")
	}
	u := models.User{}
	if models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10003, "用户名不存在")
	}

	if password == "" {
		this.ReturnJson(10004, "您的输入不完整")
	} else {
		if models.DB.Where("username = ? and password = ?", username, utils.Md5(password)).Find(&u).RecordNotFound() {
			this.ReturnJson(10005, "用户名或者密码错误")
		} else {
			this.SetSession("userinfo", &u)
			this.ReturnSuccess()
		}
	}
}

//@router /api/user/smslogin [*]
func (this *Controller) SmsLoginSubmit() {
	username := this.GetString("username")
	if username == "" {
		this.ReturnJson(10001, "用户名不能为空")
		return
	}

	if !utils.MobileRegexp(username) {
		this.ReturnJson(10002, "请输入正确的手机号码")
	}
	u := models.User{}
	if models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10003, "用户名不存在")
	}

	this.VerityCode(username)
	this.SetSession("userinfo", &u)
	this.ReturnSuccess()
}

//@router /api/user/register [*]
func (this *Controller) RegisterSubmit() {
	username := this.GetString("username")
	password1 := this.GetString("password1")
	password2 := this.GetString("password2")
	if !utils.MobileRegexp(username) {
		this.ReturnJson(10001, "请输入正确的手机号码")
	}
	if password1 == "" {
		this.ReturnJson(10002, "密码不能为空")
	}
	if password1 != password2 {
		this.ReturnJson(10002, "两次密码不一致")
	}

	//验证码验证
	this.VerityCode(username)

	u := models.User{}
	if !models.DB.Where("username = ?", username).First(&u).RecordNotFound() {
		this.ReturnJson(10005, "该用户已经注册")
	}
	beego.Debug(username, password1)
	beego.Debug(u)
	u.Username = username
	u.Password = utils.Md5(password1)
	u.Prioty = 3

	if err := models.DB.Create(&u).Error; err != nil {
		beego.Debug(err)
		this.ReturnJson(10006, "注册用户失败")
	}

	this.ReturnSuccess()
}

//@router /api/user/logout [*]
func (this *Controller) Logout() {
	this.DelSession("userinfo")
	this.Redirect("/",302)
}

//@router /api/user/send/sms [post]
func (this *Controller) SendSms() {
	username := this.GetString("username")
	if username == "" {
		this.ReturnJson(10001, "手机号不能为空")
		return
	}
	if !utils.MobileRegexp(username) {
		this.ReturnJson(10001, "请输入正确的手机号码")
		return
	}
	code := string(utils.Krand(6, 0))
	this.SetSession(username, code)
	go utils.SendMsg(username, code)
	go func(this *Controller,username string){
		time.Sleep(300*time.Second)
		this.DelSession(username)
	}(this,username)
	this.ReturnSuccess()
}

//@router /api/user/send/rsms [post]
func (this *Controller) SendrSms() {
	username := this.GetString("username")
	if username == "" {
		this.ReturnJson(10001, "手机号不能为空")
		return
	}
	if !utils.MobileRegexp(username) {
		this.ReturnJson(10001, "请输入正确的手机号码")
		return
	}
	this.CaptchaInterceptor()
	code := string(utils.Krand(6, 0))
	this.SetSession(username, code)
	go utils.SendMsg(username, code)
	this.ReturnSuccess()
}
