package user

import (
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/utils"
	"io/ioutil"
)

type UserController struct {
	controllers.BaseController
}

//@router /user/data [get]
func (this *UserController)UserData(){
	userinfo:=this.Userinfo
	this.ReturnSuccess("userinfo",userinfo)
}

//@router /user/usrn_update [post]
func (this *UserController)UsrnUpdate(){
	userinfo:=this.Userinfo
	username:=this.GetString("username")
	code:=this.GetString("code")
	if !this.VerityCode(code){
		beego.Debug("验证码错误")
		this.ReturnJson(1,"验证码错误")
	}
	userinfo.Username=username
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新手机失败",err)
		this.ReturnJson(1,"更新手机失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /user/pwd_update [post]
func (this *UserController)PwdUpdate(){
	userinfo:=this.Userinfo
	old_password:=this.GetString("old_password")
	new_password:=this.GetString("new_password")
	md5_old_pwd:=utils.Md5(old_password)
	old_password=""
	md5_new_pwd:=utils.Md5(new_password)
	new_password=""
	if md5_old_pwd!=this.Userinfo.Password{
		beego.Debug("旧密码错误")
		this.ReturnJson(1,"旧密码错误")
	}
	userinfo.Password=md5_new_pwd
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新密码失败",err)
		this.ReturnJson(1,"更新密码失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /user/img_update [post]
func (this *UserController)ImgUpdate(){
	userinfo:=this.Userinfo
	_, fileHeader, err := this.GetFile("imgFile")
	if err != nil {
		beego.Debug("get file error :", err)
		this.ReturnJson(1,"get file error :"+err.Error())
	}
	file, err := fileHeader.Open()
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Debug("读取文件流失败",err)
		this.ReturnJson(1,"读取文件流失败",err.Error())
	}

	url, err := utils.UploadFile(fileHeader.Filename, data)
	if err != nil {
		beego.Debug("上传文件失败",err)
		this.ReturnJson(1,"上传文件失败"+err.Error())
	}
	userinfo.Icon=url
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新头像失败",err)
		this.ReturnJson(1,"更新头像失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /user/sex_update [post]
func (this *UserController)SexUpdate(){
	userinfo:=this.Userinfo
	sex,_:=this.GetInt("sex")
	userinfo.Sex=sex
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新性别失败",err)
		this.ReturnJson(1,"更新性别失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /user/qq_update [post]
func (this *UserController)QQUpdate(){
	userinfo:=this.Userinfo
	qq:=this.GetString("qq")
	userinfo.QQ=qq
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新qq失败",err)
		this.ReturnJson(1,"更新qq失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /user/email_update [post]
func (this *UserController)EmailUpdate(){
	userinfo:=this.Userinfo
	email:=this.GetString("email")
	userinfo.Email=email
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新qq失败",err)
		this.ReturnJson(1,"更新qq失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

//@router /user/sign_update [post]
func (this *UserController)SignUpdate(){
	userinfo:=this.Userinfo
	sign:=this.GetString("sign")
	userinfo.Sign=sign
	if err:=models.DB.Save(&userinfo).Error;err!=nil{
		beego.Debug("更新qq失败",err)
		this.ReturnJson(1,"更新qq失败"+err.Error())
	}
	this.ReturnSuccess()
	return
}

/*
权限：
	用户不能登陆后台
	普通管理员有部分删（删除用户，删除栏目，删除资源，不能删除管理员）权限，有部分增（增加资源，增加用户，不能增加管理员，增加栏目）
修（修改资源，修改用户，不能修改管理员，修改栏目）权限，有全部查权限
	普通管理员对资源 用户 栏目 是满权限的     对管理员是无权限的   不能修改权限
	root管理员：杀人放火，不所不能
 */
//@router /article/list [get]
func (this *UserController)ArticleList(){
	userinfo:=this.Userinfo
	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per,_ = beego.AppConfig.Int("per")
	}

	qs := models.DB.Model(models.Article{})

	count := 0
	qs.Count(&count)
	articles:=[]models.Article{}
	if err:=qs.Where("user_id = ?",userinfo.ID).Limit(per).Offset((page - 1) * per).Order("id desc").Find(articles).Error;err!=nil{
		beego.Debug("读取文章列表错误"+err.Error())
		this.ReturnJson(1,"读取文章列表错误"+err.Error())
	}
	this.ReturnSuccess("articles", articles, "page", page, "count", count, "per", per)
}

//@router /article [get]
func (this *UserController)Article(){
	article:=models.Article{}
	if id,err:=this.GetByID(&article);err!=nil{
		beego.Debug("读取文章错误",err)
		this.ReturnJson(1,"读取文章错误"+err.Error())
	}else{
		beego.Debug("文章——",id)
	}
	beego.Debug(article)
	this.ReturnSuccess("article",article)
}

//@router /article/submit [post]
func (this *UserController)ArticleSubmit(){
	userinfo:=this.Userinfo
	content:=this.GetString("content")
	article:=models.Article{}
	article.UserID=userinfo.ID
	article.Content=content
	if err:=models.DB.Create(&article).Error;err!=nil{
		beego.Debug("存文章失败",err)
		this.ReturnJson(1,"存文章失败"+err.Error())
	}
	this.ReturnSuccess()
}

//@router /article/update [post]
func (this *UserController)ArticleUpdate(){
	content:=this.GetString("content")
	article:=models.Article{}
	if id,err:=this.GetByID(&article);err!=nil{
		beego.Debug("通过ID获取文章失败",err)
		this.ReturnJson(1,"通过ID获取文章失败"+err.Error())
	}else{
		beego.Debug("文章——",id)
	}
	beego.Debug(article)
	if article.Status==0{
		beego.Debug("该文章已审核通过，不可修改")
		this.ReturnJson(1,"该文章已审核通过，不可修改")
	}
	article.Content=content
	if err:=models.DB.Save(&article).Error;err!=nil{
		beego.Debug("更新文章失败",err)
		this.ReturnJson(1,"更新文章失败"+err.Error())
	}
	this.ReturnSuccess()
}
