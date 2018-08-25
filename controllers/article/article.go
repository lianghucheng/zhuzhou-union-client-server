package article

import (
	"strconv"
	"zhuzhou-union-client-server/models"
)

//@router /article/:id [*]
func (this *Controller) ArticleDetail() {

	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var article models.Article
	if err := models.DB.Preload("Category").Where("id = ?", id).
		First(&article).Error; err != nil {
		this.Abort("404")
		return
	}

	var categories []models.Category
	if article.Category != nil {
		models.DB.Where("higher_id=?", article.Category.HigherID).Find(&categories)
	}
	this.Data["categories"] = categories

	var recommend []models.Article
	models.DB.Select("id,cover,summary,title,author,created_at").
		Where("category_id=?", article.CategoryID).Order("read_num desc").Limit(6).Find(&recommend)

	this.Data["recommend"] = recommend

	this.Data["article"] = article
	this.TplName = "article/article.html"
}
