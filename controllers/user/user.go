package user

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
)

type Controller struct {
	controllers.Common
}

//@router /user/center [*]
func (this *Controller) UserCenter(){
	this.TplName="user/user.html"
}

//@router /api/user/data [get]
func (this *Controller) UserData() {
	userinfo := this.Userinfo
	this.ReturnSuccess("userinfo", userinfo)
}

//@router /api/user/usrn_update [post]
func (this *Controller) UsrnUpdate() {
	userinfo := this.Userinfo
	username := this.GetString("username")

	//this.VerityCode(username)

	if !utils.MobileRegexp(username) {
		this.ReturnJson(10002, "请输入正确的手机号码")
	}
	userinfo.Username = username
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新手机失败", err)
		this.ReturnJson(1, "更新手机失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/user/pwd_update [post]
func (this *Controller) PwdUpdate() {
	userinfo := this.Userinfo
	old_password := this.GetString("old_password")
	new_password := this.GetString("new_password")
	md5_old_pwd := utils.Md5(old_password)
	old_password = ""
	md5_new_pwd := utils.Md5(new_password)
	new_password = ""
	if md5_old_pwd != this.Userinfo.Password {
		beego.Debug("旧密码错误")
		this.ReturnJson(1, "旧密码错误")
	}
	userinfo.Password = md5_new_pwd
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新密码失败", err)
		this.ReturnJson(1, "更新密码失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/user/pwd_find [post]
func (this *Controller) PwdFind() {
	username := this.GetString("username")
	//this.VerityCode(username)
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

//@router /api/user/img_update [post]
func (this *Controller) ImgUpdate() {
	userinfo := this.Userinfo
	_, fileHeader, err := this.GetFile("imgFile")
	if err != nil {
		beego.Debug("get file error :", err)
		this.ReturnJson(1, "get file error :"+err.Error())
	}
	file, err := fileHeader.Open()
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Debug("读取文件流失败", err)
		this.ReturnJson(1, "读取文件流失败", err.Error())
	}

	url, err := utils.UploadFile(fileHeader.Filename, data)
	if err != nil {
		beego.Debug("上传文件失败", err)
		this.ReturnJson(1, "上传文件失败"+err.Error())
	}
	userinfo.Icon = url
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新头像失败", err)
		this.ReturnJson(1, "更新头像失败"+err.Error())
	}
	this.ReturnSuccess()
	return
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

//@router /api/user/qq_update [post]
func (this *Controller) QQUpdate() {
	userinfo := this.Userinfo
	qq := this.GetString("qq")
	userinfo.QQ = qq
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新qq失败", err)
		this.ReturnJson(1, "更新qq失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /api/user/email_update [post]
func (this *Controller) EmailUpdate() {
	userinfo := this.Userinfo
	email := this.GetString("email")
	userinfo.Email = email
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新qq失败", err)
		this.ReturnJson(1, "更新qq失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /api/user/sign_update [post]
func (this *Controller) SignUpdate() {
	userinfo := this.Userinfo
	sign := this.GetString("sign")
	userinfo.Sign = sign
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("更新qq失败", err)
		this.ReturnJson(1, "更新qq失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /api/user/submit_detail [post]
func (this *Controller) SubmitDetail() {
	userinfo := this.Userinfo
	sex, _ := this.GetInt("sex")
	qq := this.GetString("qq")
	email := this.GetString("email")
	sign := this.GetString("")
	name := this.GetString("name")
	_, fileHeader, err := this.GetFile("imgFile")
	if err != nil {
		beego.Debug("get file error :", err)
		this.ReturnJson(1, "get file error :"+err.Error())
	}

	file, err := fileHeader.Open()
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Debug("读取文件流失败", err)
		this.ReturnJson(1, "读取文件流失败", err.Error())
	}

	url, err := utils.UploadFile(fileHeader.Filename, data)
	if err != nil {
		beego.Debug("上传文件失败", err)
		this.ReturnJson(1, "上传文件失败"+err.Error())
	}
	userinfo.Icon = url
	userinfo.Sex = sex
	userinfo.QQ = qq
	userinfo.Email = email
	userinfo.Sign = sign
	userinfo.Name = name
	if err := models.DB.Save(&userinfo).Error; err != nil {
		beego.Debug("保存失败", err)
		this.ReturnJson(1, "保存失败"+err.Error())
	}
	this.ReturnSuccess()
}

/*
权限：
	用户不能登陆后台
	普通管理员有部分删（删除用户，删除栏目，删除资源，不能删除管理员）权限，有部分增（增加资源，增加用户，不能增加管理员，增加栏目）
修（修改资源，修改用户，不能修改管理员，修改栏目）权限，有全部查权限
	普通管理员对资源 用户 栏目 是满权限的     对管理员是无权限的   不能修改权限
	root管理员：杀人放火，不所不能
*/
//@router /api/article/list [get]
func (this *Controller) ArticleList() {
	userinfo := this.Userinfo
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
	userinfo := this.Userinfo
	content := this.GetString("content")
	article := models.Article{}
	article.UserID = userinfo.ID
	article.Content = content
	if err := models.DB.Create(&article).Error; err != nil {
		beego.Debug("存文章失败", err)
		this.ReturnJson(1, "存文章失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /api/article/update [post]
func (this *Controller) ArticleUpdate() {
	content := this.GetString("content")
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
