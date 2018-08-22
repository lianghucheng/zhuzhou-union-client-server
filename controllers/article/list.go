package article

import (
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
	models.DB.Where("id=?", id).Find(&category)

	this.TplName = "article/category.html"
}
