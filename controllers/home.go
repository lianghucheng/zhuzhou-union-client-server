package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
	"github.com/lexkong/log"

)

type HomeController struct {
	Common
}



//@router /	[*]
func (this *HomeController) Index() {
	var homes []*models.Home

	var rotation []*models.Rotation

	var indexPer int

	indexPer, _ = beego.AppConfig.Int("indexPer")
	if indexPer == 0 {
		indexPer = 5
	}


	//首页文章
	if err := models.DB.
		Preload("Category").
		Find(&homes).Error; err != nil {
		log.Error("获取首页表数据错误", err)
		this.Abort("500")
		return
	}

	if err := models.DB.
		Find(&rotation).Error; err != nil {
		beego.Error("获取首页轮播图错误", err)
	}

	outputIndex := make([]map[string]interface{}, 0)

	//首页子分类
	for _, h := range homes {
		var articles1 []*models.Article
		subCatesM := make([]map[string]interface{}, 0)
		a := make(map[string]interface{})
		var subCategorys []*models.Category
		if (h.Position == 1 && h.Layout == 1) {
			if err := models.DB.
				Where("higher_id = ?", h.CategoryID).
				Find(&subCategorys).Error; err != nil {
				beego.Error("读取首页子分类错误", err)
				this.Abort("500")
				return
			}

			for _, subCategory := range subCategorys {
				var articles []*models.Article
				subCateM := make(map[string]interface{})

				if err := models.DB.
					Where("category_id = ? and status =?", subCategory.ID, 1).Limit(indexPer).
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

		if h.Position == 13 || h.Position == 14 {
			indexPer = 5
		}

		if h.Position == 15 || h.Position == 16 {
			indexPer = 10
		}
		if err := models.DB.
			Where("category_id = ? and status = ?", h.CategoryID, 1).Limit(indexPer).
			Find(&articles1).Error; err != nil {
			this.Abort("500")
			return
		}

		a["Articles"] = articles1
		a["Home"] = h
		a["SubCates"] = subCatesM
		outputIndex = append(outputIndex, a)
	}
	//首页底部图片链接
	var imageLinks []*models.ImageLinks
	if err := models.DB.Limit(5).Find(&imageLinks).Error; err != nil {
		beego.Error("获取首页图片链接错误", err)
	}

	//首页底部下拉框链接
	var boxLinks []*models.BoxLinks

	if err := models.DB.Find(&boxLinks).Error; err != nil {
		beego.Error("获取首页下拉链接错误", err)
	}

	this.Data["rotation"] = rotation
	this.Data["imageLinks"] = imageLinks
	this.Data["homes"] = homes
	this.Data["output"] = outputIndex
	this.Data["boxLinks"] = boxLinks
	this.TplName = "index.html"
}
