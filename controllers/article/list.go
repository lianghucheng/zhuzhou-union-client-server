package article

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
)

type Controller struct {
	controllers.Common
}

//@router /category/id [*]
func (this *Controller) List() {
	var cateId int
	this.Ctx.Input.Bind(&cateId, "id")
	//获取该分类id的子分类列表
	var subCategory []*models.Category

	if err := models.DB.
		Where("higher_id = ?", cateId).
		Find(&subCategory).Error; err != nil {
		beego.Error("获取子分类错误", err)
	}

	//已经设置好的某一分类的文章列表
	var specialCate models.Category
	if err := models.DB.
		Where("special = ?", 1).
		First(&specialCate).Error; err != nil {
		beego.Error("获取文章列表特殊模块列表分类错误", err)
	}

	var specialArticle []*models.Article
	if err := models.DB.
		Where("category_id = ?", specialCate.ID).
		Find(&specialArticle).Error; err != nil {
		beego.Error("获取特殊文章列表错误")
	}
	this.Data["subCategory"] = subCategory
	this.Data["specialCate"] = specialCate
	this.Data["specialArticle"] = specialArticle
	this.TplName = ""
}

//@router /article/list/id [post]
func (this *Controller) ArticleList() {
	var cateId int
	var page int
	var per int
	var count int
	this.Ctx.Input.Bind(&cateId, "id")
	this.Ctx.Input.Bind(&page, "p")

	if page == 0 {
		page = 1
	}

	//获取该分类分页的文章列表，
	var articles []*models.Article

	per, _ = beego.AppConfig.Int("per")
	if err := models.DB.
		Where("category_id = ?", cateId).
		Limit(per).Offset((page - 1) * per).
		Find(&articles).Count(&count).Error; err != nil {
		beego.Error("获取分类文章列表错误", err)
	}

	this.ReturnSuccess("articles", articles, "per", per, "page", page, "count", count)

}
