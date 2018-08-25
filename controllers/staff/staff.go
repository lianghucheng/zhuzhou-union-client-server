package staff

import (
	"zhuzhou-union-client-server/models"
	"strconv"
	"zhuzhou-union-client-server/controllers"
	"github.com/astaxie/beego"
)

type StaffShowController struct {
	controllers.Common
}

//@router /api/staff/show/submit [POST]
func (this *StaffShowController) Submit() {
	title := this.GetString("title")
	content := this.GetString("content")
	category, err := this.GetInt("category")
	if err != nil || category <= 0 || category < 0 || category > 3 {
		this.ReturnJson(10001, "该分类不存在")
		return
	}

	var staffShow models.StaffShow

	staffShow.Category = category
	staffShow.Title = title
	staffShow.Content = content
	staffShow.UserID = uint(this.UserID)

	if err := models.DB.Create(&staffShow).Error; err != nil {
		this.ReturnJson(10002, "添加作品失败")
		return
	}
	this.ReturnSuccess()
}

//@router /api/staff/show/list [*]
func (this *StaffShowController) List() {
	page, err := this.GetInt("p")

	if err != nil || page == 0 {
		page = 1
	}
	per, _ := beego.AppConfig.Int("per")

	var staffShows []*models.StaffShow
	if err := models.DB.Offset((page - 1) * per).Limit(per).
		Find(&staffShows).Error; err != nil {
		this.ReturnJson(10001, "获取数据错误")
		return
	}

	this.Data["StaffShows"] = staffShows
	this.TplName = "staff/staff.html"

}

//@router /api/staff/show/update [POST]
func (this *StaffShowController) Update() {

}

//@router /api/staff/show/delete/id [POST]
func (this *StaffShowController) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if id == 0 {
		this.ReturnJson(10001, "用户不存在")
		return
	}
	var staffShow models.StaffShow

	if err := models.DB.Where("id = ?", id).
		Delete(&staffShow).Error; err != nil {
		this.ReturnJson(10002, "用户不存在")
		return
	}

	this.ReturnSuccess()
}
