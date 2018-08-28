package user

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
	"io/ioutil"
)

type Controller struct {
	controllers.Common
}

//@router /user/center [*]
func (this *Controller) UserCenter() {
	this.CheckLogin()
	this.TplName = "user/user.html"
}

//@router /api/user/data [get]
func (this *Controller) UserData() {
	this.CheckLogin()
	userinfo := this.GetSession("userinfo").(*models.User)

	models.DB.Where("id  = ?", userinfo.ID).
		First(&userinfo)

	this.ReturnSuccess("userinfo", userinfo)
}

//@router /api/user/usrn_update [post]
func (this *Controller) UsrnUpdate() {
	this.CheckLogin()
	userinfo := this.GetSession("userinfo").(*models.User)
	username := this.GetString("username")

	this.VerityCode(username)

	if !utils.MobileRegexp(username) {
		this.ReturnJson(10002, "请输入正确的手机号码")
	}
	userinfo.Username = username
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新手机失败", err)
		this.ReturnJson(1, "更新手机失败"+err.Error())
	}
	this.ReturnSuccess("userinfo", userinfo)
}

//@router /api/user/pwd_update [post]
func (this *Controller) PwdUpdate() {
	this.CheckLogin()
	userinfo := this.GetSession("userinfo").(*models.User)
	old_password := this.GetString("old_password")
	new_password := this.GetString("new_password")
	md5_old_pwd := utils.Md5(old_password)
	old_password = ""
	md5_new_pwd := utils.Md5(new_password)
	new_password = ""
	if md5_old_pwd != userinfo.Password {
		beego.Debug("旧密码错误")
		this.ReturnJson(1, "旧密码错误")
	}
	userinfo.Password = md5_new_pwd
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新密码失败", err)
		this.ReturnJson(1, "更新密码失败"+err.Error())
	}
	this.ReturnSuccess("userinfo", userinfo)
}

//@router /api/user/pwd_find [post]
func (this *Controller) PwdFind() {
	username := this.GetString("username")
	this.VerityCode(username)
	new_pwd := this.GetString("new_pwd")
	md5_pwd := utils.Md5(new_pwd)
	new_pwd = ""
	userinfo := models.User{}
	if err := models.
		DB.
		Where("username = ?", username).
		Find(&userinfo).
		Error; err != nil {
		beego.Debug("找回密码--读取用户失败" + err.Error())
		this.ReturnJson(1, "找回密码--读取用户失败"+err.Error())
	}
	userinfo.Password = md5_pwd
	if err := models.
		DB.
		Save(&userinfo).
		Error; err != nil {
		beego.Debug("找回密码--保存密码失败" + err.Error())
		this.ReturnJson(1, "找回密码--保存密码失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/article/list [get]
func (this *Controller) ArticleList() {
	this.CheckLogin()
	userinfo := this.GetSession("userinfo").(*models.User)
	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per, _ = beego.AppConfig.Int("per")
	}

	qs := models.DB.Model(models.Article{})

	sum := 0
	qs.Count(&sum)
	count := 0
	qs.Where("user_id = ?", userinfo.ID).Count(&count)
	articles := []models.Article{}
	if err := qs.Where("user_id = ?", userinfo.ID).Limit(per).Offset((page - 1) * per).Order("id desc").Find(&articles).Error; err != nil {
		beego.Debug("读取文章列表错误" + err.Error())
		this.ReturnJson(1, "读取文章列表错误"+err.Error())
	}
	this.ReturnSuccess("articles", articles, "page", page, "sum", sum, "count", count, "per", per)
}

//@router /api/article [get]
func (this *Controller) Article() {
	article := models.Article{}
	if id, err := this.GetByID(&article); err != nil {
		beego.Debug("读取文章错误", err)
		this.ReturnJson(1, "读取文章错误"+err.Error())
	} else {
		beego.Debug("文章——", id)
	}
	beego.Debug(article)
	this.ReturnSuccess("article", article)
}

//@router /api/article/submit [post]
func (this *Controller) ArticleSubmit() {
	this.CheckLogin()
	userinfo := this.GetSession("userinfo").(*models.User)
	content := this.GetString("content")
	title := this.GetString("title")
	cid, _ := this.GetInt("cid")
	article := models.Article{}
	article.UserID = userinfo.ID
	article.Content = content
	article.CategoryID = uint(cid)
	article.Title = title
	if err := models.DB.Create(&article).Error; err != nil {
		beego.Debug("存文章失败", err)
		this.ReturnJson(1, "存文章失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/article/update [post]
func (this *Controller) ArticleUpdate() {
	content := this.GetString("content")
	cid, _ := this.GetInt("cid")
	article := models.Article{}
	if id, err := this.GetByID(&article); err != nil {
		beego.Debug("通过ID获取文章失败", err)
		this.ReturnJson(1, "通过ID获取文章失败"+err.Error())
	} else {
		beego.Debug("文章——", id)
	}
	beego.Debug(article)
	if article.Status == 1 {
		beego.Debug("该文章已审核通过，不可修改")
		this.ReturnJson(1, "该文章已审核通过，不可修改")
	}
	article.Content = content
	article.CategoryID = uint(cid)
	if err := models.DB.Save(&article).Error; err != nil {
		beego.Debug("更新文章失败", err)
		this.ReturnJson(1, "更新文章失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/user/search [get]
func (this *Controller) Search() {
	userinfo := this.Userinfo
	searcher := this.GetString("searcher")
	s := "%" + searcher + "%"

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per, _ = beego.AppConfig.Int("per")
	}

	qs := models.DB.Model(models.Article{})

	sum := 0
	qs.Count(&sum)
	count := 0
	qs.Where("user_id = ?", userinfo.ID).Count(&count)
	articles := []models.Article{}
	if err := qs.Where("user_id = ? and (title LIKE ? or content LIKE ?)", userinfo.ID, s, s).Limit(per).Offset((page - 1) * per).Order("id desc").Find(&articles).Error; err != nil {
		beego.Error("查找失败", err)
		this.ReturnJson(1, "查找失败"+err.Error())
	}
	this.ReturnSuccess("articles", articles, "page", page, "sum", sum, "count", count, "per", per)
}

//@router /api/user/category [get]
func (this *Controller) GetCate() {
	categorys := []models.Category{}
	if err := models.DB.Where("is_submission = ?", 1).Find(&categorys).Error; err != nil {
		beego.Error("取分类失败", err)
		this.ReturnJson(10001, "取分类失败"+err.Error())
	}
	this.ReturnSuccess("categories", categorys)
}

//@router /api/user/update [post]
func (this *Controller) UpdateUser() {
	var userinfo models.User

	nickname := this.GetString("nickname")
	sex, _ := this.GetInt("sex")
	id, _ := this.GetInt("id")

	models.DB.Where("id = ?", id).First(&userinfo)

	userinfo.Sex = sex
	userinfo.NickName = nickname
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新姓名失败", err)
		this.ReturnJson(10001, "更新姓名失败"+err.Error())
		return
	}
	this.ReturnSuccess()
	return
}

//@router /api/user/img/upload [post]
func (this *Controller) ImgUpload() {
	file, header, _ := this.GetFile("file")

	fileByte, _ := ioutil.ReadAll(file)
	url, err := utils.UploadFile(header.Filename, fileByte)
	if err != nil {
		this.ReturnJson(10001, "上传失败")
		return
	}
	beego.Debug(url)
	m := make(map[string]interface{})
	m["url"] = url
	m["status"] = 10000
	m["message"] = "success"
	this.Data["json"] = m
	this.ServeJSON()
	return
}
