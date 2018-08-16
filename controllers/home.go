package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
	"github.com/lexkong/log"
	"fmt"
)

type HomeController struct {
	beego.Controller
}

//@router /	[*]
func (this *HomeController) Index() {
	var homes []*models.Home

	var rotation []*models.Rotation
	var Menus []*models.Menu
	var indexPer int

	indexPer, _ = beego.AppConfig.Int("indexPer")
	if indexPer == 0 {
		indexPer = 5
	}

	if err := models.DB.Preload("Category").Find(&Menus).Error; err != nil {
		beego.Error("查询菜单错误", err)
	}

	var outputMenus []models.Menu
	for _, menu := range Menus {
		var outputMenu models.Menu
		if menu.CategoryID != 0 {
			category := menu.Category

			var categoryMenu models.Menu
			categoryMenu.Name = category.Name
			categoryMenu.URL = "/category/" + fmt.Sprintf("%s", category.ID)

			var subCategorys []*models.Category
			if err := models.DB.
				Where("higher_id = ?", category.ID).
				Find(&subCategorys).Error; err != nil {
				beego.Error("查询子菜单错误")
			}

			for _, subCategory := range subCategorys {
				var itemMenu models.Menu
				itemMenu.Name = subCategory.Name
				itemMenu.URL = "/category/" + fmt.Sprintf("%s", subCategory.ID)
				categoryMenu.Menus = append(categoryMenu.Menus, itemMenu)
			}
			outputMenu = categoryMenu

		} else {
			var notCategoryMenu models.Menu
			notCategoryMenu.Name = menu.Name
			notCategoryMenu.URL = menu.URL
			outputMenu = notCategoryMenu
		}
		outputMenus = append(outputMenus, outputMenu)
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
	if err := models.DB.Find(&imageLinks).Error; err != nil {
		beego.Error("获取首页图片链接错误", err)
	}

	//首页底部下拉框链接
	var boxImages []*models.BoxLinks

	if err := models.DB.Find(&boxImages).Error; err != nil {
		beego.Error("获取首页下拉链接错误", err)
	}

	this.Data["outputMenus"] = outputMenus
	this.Data["rotation"] = rotation
	this.Data["imageLinks"] = imageLinks
	this.Data["homes"] = homes
	this.Data["output"] = outputIndex
	this.TplName = "index.html"
}
