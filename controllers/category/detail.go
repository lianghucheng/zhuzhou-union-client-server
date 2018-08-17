package category

import (
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego"
)

//@router /article/detail/id [*]
func (this *CategoryController) ArticleDetail() {
	articleId, _ := this.GetInt("id")

	var article models.Article
	if err := models.DB.Where("id = ?", articleId).
		First(&article).Error; err != nil {
		beego.Error("没有此文章")

	}
	var category models.Category

	if err := models.DB.Where("id = ?", article.CategoryID).
		First(&category).Error; err != nil {
		beego.Error("没有此分类")
	}

	//这是渲染侧边栏的分类
	var subCategory []*models.Category
	//爸爸找儿子
	if category.HigherID == 0 {
		if err := models.DB.Where("higher_id = ?", category.ID).
			Find(&subCategory).Error; err != nil {
			beego.Error("没有找到该分类的二级分类")
		}
		this.Data["subCategory"] = subCategory
		this.Data["breadCrumbs"] = subCategory
		beego.Error("此分类是顶级分类,一般不会出现此种情况，一般此情况出现为后台录入时直接将文章放在顶级分类下")
	} else {
		//直接找兄弟
		if err := models.DB.Where("higher_id = ?", category.HigherID).
			Find(&subCategory).Error; err != nil {
			beego.Error("没有找到该分类的二级分类")
		}
		this.Data["subCategory"] = subCategory
		this.Data["breadCrumbs"] = subCategory
	}

	//侧边栏特殊分类的文章
	var specialCategorys []*models.Category

	if err := models.DB.Where("special = ?", 1).
		Find(&specialCategorys).Error; err != nil {
		beego.Error("没有找到特殊分类")
	}

	outputSpecail := make([]map[string]interface{}, 0)

	for _, specialCategory := range specialCategorys {
		var articles []*models.Article
		output := make(map[string]interface{})
		if err := models.DB.Where("category_id = ?", specialCategory.ID).
			Find(&articles).Error; err != nil {
			beego.Error("没有找到该分类的文章")
		}

		output["articles"] = articles
		output["category"] = specialCategory
		outputSpecail = append(outputSpecail, output)
	}

	this.Data["special"] = outputSpecail
	this.Data["article"] = article
	this.TplName = "category/detail.html"

}
