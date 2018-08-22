package article

import (
	"strconv"
	"zhuzhou-union-client-server/models"
)

//@router /article/:id [*]
func (this *Controller) ArticleDetail() {

	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var article models.Article
	if err := models.DB.Where("id = ?", id).
		First(&article).Error; err != nil {
		this.Abort("404")
		return
	}

	this.Data["article"] = article
	this.TplName = "article/article.html"
}
