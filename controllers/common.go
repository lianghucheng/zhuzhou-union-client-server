package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
	"fmt"
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

	var code models.QrCode

	user := models.User{
		Username: "test",
	}
	var Menus []*models.Menu
	if err := models.DB.Preload("Category").Where("higher_id = ?", 0).Find(&Menus).Error; err != nil {
		beego.Error("查询菜单错误", err)
	}
	var outputMenus []models.Menu
	for _, menu := range Menus {
		var outputMenu models.Menu
		if menu.CategoryID != 0 {
			category := menu.Category

			var categoryMenu models.Menu
			categoryMenu.Name = menu.Name
			categoryMenu.URL = "/category/" + fmt.Sprintf("%d", category.ID)
			categoryMenu.Sequence = menu.Sequence
			var subCategorys []*models.Category
			if err := models.DB.
				Where("higher_id = ?", category.ID).
				Find(&subCategorys).Error; err != nil {
				beego.Error("查询子菜单错误")
			}

			for _, subCategory := range subCategorys {
				var itemMenu models.Menu
				itemMenu.Name = subCategory.Name
				itemMenu.URL = "/category/" + fmt.Sprintf("%d", subCategory.ID)

				categoryMenu.Menus = append(categoryMenu.Menus, itemMenu)
			}
			outputMenu = categoryMenu
		} else {
			var notCategoryMenu models.Menu

			notCategoryMenu.Name = menu.Name
			notCategoryMenu.URL = menu.URL
			notCategoryMenu.Sequence = menu.Sequence
			outputMenu = notCategoryMenu
		}
		outputMenus = append(outputMenus, outputMenu)
	}
	if err := models.DB.First(&code).Error; err != nil {
		beego.Error("没有发现二维码")
	}
	this.SetSession("userinfo", user)
	user, ok := this.GetSession("userinfo").(models.User)
	if ok {
		this.Data["user"] = user
	}
	this.Data["Code"] = code
	this.Data["outputMenus"] = outputMenus
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
