package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
	"github.com/dchest/captcha"
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

type OutputMenu struct {
	Name  string
	URL   string
	Menus []OutputMenu
}

func (this *Common) initFooter() {

	//首页底部图片链接
	var imageLinks []*models.ImageLinks
	models.DB.Find(&imageLinks)

	//首页底部下拉框链接
	var boxLinks []*models.BoxLinks
	models.DB.Find(&boxLinks)

	this.Data["boxLinks"] = boxLinks
	this.Data["imageLinks"] = imageLinks
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

	var outputs []OutputMenu

	for _, menu := range menus {

		var output OutputMenu
		output.Name = menu.Name
		output.URL = menu.URL

		if menu.CategoryID != 0 {
			output.URL = "/category/" + fmt.Sprintf("%d", menu.CategoryID)
			var subCategories []models.Category
			models.DB.Where("higher_id=?", menu.CategoryID).Find(&subCategories)
			for _, childCategory := range subCategories {
				var child OutputMenu
				child.URL = "/category/" + fmt.Sprintf("%d", childCategory.ID)
				child.Name = childCategory.Name
				output.Menus = append(output.Menus, child)
			}
		} else {
			for _, childMenu := range menu.Menus {
				var child OutputMenu
				child.URL = childMenu.URL
				child.Name = childMenu.Name
				output.Menus = append(output.Menus, child)
			}

		}

		outputs = append(outputs, output)
	}

	this.Data["outputMenus"] = outputs
}

func (this *Common) Prepare() {

	this.initMenu()
	this.initFooter()
	if this.IsLogin() {
		this.Data["User"] = this.GetSessionUser()
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
			this.ReturnJson(10003, "验证码输入错误")
		}
	} else {
		this.ReturnJson(10004, "验证码输入错误")
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

func (ctr *Common) CaptchaInterceptor() {
	code := ctr.GetString("captcha_code")
	captchaId := ctr.GetString("captcha_id")
	if !captcha.VerifyString(captchaId, code) {
		ctr.ReturnJson(10401, "captcha verify error")
		ctr.StopRun()
		return
	}
}
