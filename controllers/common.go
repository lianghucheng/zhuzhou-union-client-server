package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
)

type Common struct {
	beego.Controller
	UserID   int64
	Token    string
	Userinfo *models.User
}

type CommonController struct {
	Common
}

func (this *Common) initMenu() {
	var menus []*models.Menu
	if err := models.DB.
		Preload("Menus").
		Preload("Category").
		Where("higher_id = ?", 0).
		Order("sequence asc").Find(&menus).Error; err != nil {
		beego.Error("查询菜单错误", err)
	}
	this.Data["outputMenus"] = menus
}

func (this *Common) Prepare() {

	this.initMenu()

	if this.IsLogin() {
		this.Data["user"] = this.GetSessionUser()
	}

	var code models.QrCode
	if err := models.DB.First(&code).Error; err != nil {
		beego.Error("没有发现二维码")
	}

	this.Data["Code"] = code

}

func (this *Common) UserFilter() {
}

func (this *Common) GetByID(obj interface{}) (int64, error) {
	id, _ := this.GetInt64("id")
	return id, models.DB.Where("id=?", id).First(obj).Error
}

func (this *Common) VerityCode(username string) {
	code := this.GetString("code")
	if code == "" {
		this.DelSession(username)
		this.ReturnJson(10003, "验证码不能为空")
	}

	if local_code, ok := this.GetSession(username).(string); ok {
		if local_code == code {
			this.DelSession(username)
		} else {
			this.DelSession(username)
			this.ReturnJson(10003, "验证码无效")
		}
	} else {
		this.DelSession(username)
		this.ReturnJson(10004, "验证码无效")
	}
}

func (this *Common) ReturnJson(status int, message string, args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = status
	result["message"] = message

	this.GetString("")

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

func (this *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = 10000
	result["message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

func (this *Common) DoLogin(user models.User) {
	this.SetSession("userinfo", user)
}

func (this *Common) IsLogin() bool {
	return this.GetSession("userinfo") != nil
}

func (this *Common) CheckLogin() {
	if !this.IsLogin() {
		this.Ctx.Redirect(302, "/user/login")
		this.StopRun()
		return
	}
}

func (this *Common) CheckLoginPost() {
	if !this.IsLogin() {
		this.ReturnJson(10043, "请登录")
		this.StopRun()
		return
	}
}

func (this *Common) GetSessionUser() (user *models.User) {
	return this.GetSession("userinfo").(*models.User)

}
