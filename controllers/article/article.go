package article

import (
	"strconv"
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego"
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
		if article.Category.HigherID != 0 {
			models.DB.Where("higher_id=?", article.Category.HigherID).Find(&categories)
		}
	}
	this.Data["categories"] = categories

	var recommend []models.Article
	models.DB.Select("id,cover,summary,title,author,created_at").
		Where("category_id=?", article.CategoryID).Order("read_num desc").Limit(6).Find(&recommend)

	categoryStack := make([]models.Category, 0)
	categoryStack = append(categoryStack, *article.Category)
	categoryStack = getCategoryStack(categoryStack)
	for i, j := 0, len(categoryStack)-1; i < j; i, j = i+1, j-1 {
		categoryStack[i], categoryStack[j] = categoryStack[j], categoryStack[i]
	}

	this.Data["categoryStack"] = categoryStack

	article.ReadNum = article.ReadNum + 1
	models.DB.Save(&article)
	this.Data["recommend"] = recommend
	this.Data["article"] = article
	this.TplName = "article/article.html"
}
