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
	var subCatesM []map[string]interface{}
	if err := models.DB.Preload("Category").Preload("IndexArticle").Find(&homes).Error; err != nil {
		log.Error("获取首页表数据错误", err)
		this.Abort("500")
		return
	}
	output := make([]map[string]interface{}, 0)
	for _, h := range homes {
		var articles1 []*models.Article
		subCatesM = make([]map[string]interface{}, 0)
		a := make(map[string]interface{})
		var subCategorys []*models.Category
		if (h.Position == 1 && h.Layout == 1) || (h.Position == 4 && h.Layout == 1) {
			if err := models.DB.
				Where("higher_id = ?", h.CategoryID).
				Find(subCategorys).Error; err != nil {
				beego.Error("读取首页子分类错误", err)
				this.Abort("500")
				return
			}

			for _, subCategory := range subCategorys {
				var articles []*models.Article
				subCateM := make(map[string]interface{})

				if err := models.DB.
					Where("category_id = ?", subCategory.ID).
					Find(&articles).Error; err != nil {
					beego.Error("读取首页子分类错误", err)
					this.Abort("500")
					return
				}
				subCateM["ID"] = subCategory.ID
				subCateM["Name"] = subCategory.Name
				subCateM["HigherID"] = subCategory.HigherID
				subCateM["Sequence"] = subCategory.Sequence
				subCateM["CreatedAt"] = subCategory.CreatedAt
				subCateM["DeletedAt"] = subCategory.DeletedAt
				subCateM["UpdatedAt"] = subCategory.UpdatedAt
				subCateM["Category"] = subCategory.Category
				subCateM["Articles"] = articles
				subCatesM = append(subCatesM, subCateM)
			}
		}

		if err := models.DB.
			Where("category_id = ?", h.CategoryID).
			Find(&articles1).Error; err != nil {
			this.Abort("500")
			return
		}
		a["articles"] = articles1
		a["home"] = h
		output = append(output, a)
	}

	if len(subCatesM) != 0 {
		this.Data["subCatesM"] = subCatesM
	}
	this.Data["homes"] = homes
	this.Data["output"] = output
	this.TplName = "index.html"
}
