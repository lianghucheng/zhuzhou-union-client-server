package user

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
)

type Controller struct {
	controllers.Common
}

//@router /user/center [*]
func (this *Controller) UserCenter(){
	this.CheckLogin()
	this.TplName="user/user.html"
}

//@router /api/user/data [get]
func (this *Controller) UserData() {
	this.CheckLogin()
	userinfo := this.GetSession("userinfo").(*models.User)
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
	this.ReturnSuccess("userinfo",userinfo)
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
	this.ReturnSuccess("userinfo",userinfo)
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
	title:=this.GetString("title")
	cid,_:=this.GetInt("cid")
	article := models.Article{}
	article.UserID = userinfo.ID
	article.Content = content
	article.CategoryID=uint(cid)
	article.Title=title
	if err := models.DB.Create(&article).Error; err != nil {
		beego.Debug("存文章失败", err)
		this.ReturnJson(1, "存文章失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/article/update [post]
func (this *Controller) ArticleUpdate() {
	content := this.GetString("content")
	cid,_:=this.GetInt("cid")
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
	article.CategoryID =uint(cid)
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
func (this *Controller)GetCate(){
	category:=[]models.Category{}
	if err:=models.DB.Where("category = ?","4").Find(&category).Error;err!=nil{
		beego.Error("取分类失败",err)
		this.ReturnJson(10001,"取分类失败"+err.Error())
	}
	beego.Debug(category)
	this.ReturnSuccess("categories",category)
}

//@router /api/user/name_update [post]
func (this *Controller) NameUpdate() {
	userinfo := this.Userinfo
	name := this.GetString("name")
	userinfo.Name = name
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新姓名失败", err)
		this.ReturnJson(1, "更新姓名失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /api/user/sex_update [post]
func (this *Controller) SexUpdate() {
	userinfo := this.Userinfo
	sex, _ := this.GetInt("sex")
	userinfo.Sex = sex
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新性别失败", err)
		this.ReturnJson(1, "更新性别失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}