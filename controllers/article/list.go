package article

import (
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"zhuzhou-union-client-server/controllers"
	"zhuzhou-union-client-server/models"
)

type Controller struct {
	controllers.Common
}

// @router /category/spacial/:id [*]
func (this *Controller) ListSpacial() {

	var id int
	this.Ctx.Input.Bind(&id, ":id")
	var category models.Category
	models.DB.Where("id=?", id).First(&category)

	var categories []models.Category
	models.DB.Where("higher_id = ?", category.ID).Find(&categories)

	selected, _ := this.GetInt64("selected")

	var subCategory models.Category
	if err := models.DB.Where("id=?", selected).First(&subCategory).Error; err != nil {
		if len(categories) > 0 {
			subCategory = categories[0]
		}
	}

	var articles []models.Article
	models.DB.Select("id,cover,summary,title,author,created_at").
		Where("category_id=?", subCategory.ID).Limit(3).Find(&articles)

	categoryStack := make([]models.Category, 0)
	categoryStack = append(categoryStack, category)
	categoryStack = getCategoryStack(categoryStack)
	for i, j := 0, len(categoryStack)-1; i < j; i, j = i+1, j-1 {
		categoryStack[i], categoryStack[j] = categoryStack[j], categoryStack[i]
	}
	this.Data["categoryStack"] = categoryStack

	this.Data["articles"] = articles
	this.Data["subCategory"] = subCategory
	this.Data["category"] = category
	this.Data["categories"] = categories
	this.TplName = "article/category_spacial.html"
}

func getCategoryStack(categories []models.Category) []models.Category {
	category := categories[len(categories)-1]
	if category.HigherID == 0 {
		return categories
	} else {
		var c models.Category
		models.DB.Where("id=?", category.HigherID).First(&c)
		categories = append(categories, c)
		return getCategoryStack(categories)
	}
}

//@router /category/:id [*]
func (this *Controller) List() {
	var id int
	this.Ctx.Input.Bind(&id, ":id")
	var category models.Category
	models.DB.Where("id=?", id).First(&category)

	if category.Category == 2 {
		this.Redirect(fmt.Sprintf("/category/spacial/%d", id), 302)
		return
	}

	categoryStack := make([]models.Category, 0)
	categoryStack = append(categoryStack, category)
	categoryStack = getCategoryStack(categoryStack)
	for i, j := 0, len(categoryStack)-1; i < j; i, j = i+1, j-1 {
		categoryStack[i], categoryStack[j] = categoryStack[j], categoryStack[i]
	}
	this.Data["categoryStack"] = categoryStack

	var categories []models.Category
	if category.HigherID == 0 {
		models.DB.Where("higher_id = ?", category.ID).Find(&categories)
	} else {
		models.DB.Where("higher_id=?", category.HigherID).Find(&categories)
	}
	this.Data["category"] = category
	this.Data["categories"] = categories

	pers := 6
	qs := models.DB.Select("id,cover,summary,title,author,created_at").
		Model(models.Article{}).Where("category_id=?", id)
	cnt := 0
	qs.Count(&cnt)

	pager := pagination.NewPaginator(this.Ctx.Request, pers, cnt)
	qs = qs.Order("created_at desc").Limit(pers).Offset(pager.Offset())
	var articles []*models.Article
	qs.Find(&articles)

	this.Data["articles"] = articles
	this.Data["paginator"] = pager

	var recommend []models.Article
	models.DB.Select("id,cover,summary,title,author,created_at").
		Where("category_id=? or category_id=?", id, category.HigherID).Order("read_num desc").Limit(6).Find(&recommend)

	this.Data["recommend"] = recommend

	this.TplName = "article/category.html"
}
