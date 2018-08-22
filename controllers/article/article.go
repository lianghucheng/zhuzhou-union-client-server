package article

import (
	"github.com/astaxie/beego"
	"strconv"
	"zhuzhou-union-client-server/models"
)

//@router /article/:id [*]
func (this *Controller) ArticleDetail() {

	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var article models.Article
	if err := models.DB.Where("id = ?", id).
		First(&article).Error; err != nil {
		beego.Error("没有此文章")

	}

	this.Data["article"] = article
	this.TplName = "article/article.html"
}
