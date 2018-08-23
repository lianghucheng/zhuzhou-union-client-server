package article

import (
	"github.com/astaxie/beego/utils/pagination"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
)

type Controller struct {
	controllers.Common
}

//@router /category/:id [*]
func (this *Controller) List() {
	var id int
	this.Ctx.Input.Bind(&id, ":id")
	var category models.Category
	models.DB.Where("id=?", id).First(&category)
	var categories []models.Category
	models.DB.Where("higher_id=?", category.HigherID).Find(&categories)
	this.Data["category"] = category
	this.Data["categories"] = categories

	pers := 6
	qs := models.DB.Model(models.Article{}).Where("category_id=?", id)
	cnt := 0
	qs.Count(&cnt)

	pager := pagination.NewPaginator(this.Ctx.Request, pers, cnt)
	qs = qs.Order("created_at desc").Limit(pers).Offset(pager.Offset())
	var articles []*models.Article
	qs.Find(&articles)

	this.Data["articles"] = articles
	this.Data["paginator"] = pager

	this.TplName = "article/category.html"
}
