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

func (this *Common) Prepare() {
	var user models.User

	user.Name = "test"

	this.Userinfo = &user
	if this.Userinfo != nil {
		this.Data["user"] = this.Userinfo
	}
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
