package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
	"github.com/lexkong/log"
)

type HomeController struct {
	beego.Controller
}

//@router /	[*]
func (this *HomeController) Index() {
	var homes []*models.Home
	output := make([]map[string]interface{}, 0)

	if err := models.DB.Preload("Category").Preload("IndexArticle").Find(&homes).Error; err != nil {
		log.Error("首页获取数据库数据错误", err)
		this.Abort("500")
		return
	}

	for _, h := range homes {
		var articles []*models.Article
		a := make(map[string]interface{})
		models.DB.
			Where("category_id = ? and is_index = ?", h.CategoryID, 1).
			Limit(5).
			Find(&articles)
		a["articles"] = articles

		output = append(output, a)
	}

	this.TplName = "index.html"
}
