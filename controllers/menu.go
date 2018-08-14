package controllers

import (
	"zhuzhou-union-client-server/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Menu struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Menus []Menu `json:"menus"`
}

func (this *CommonController) GetMenus() {

	var menu1 models.Menu
	menu1.Name = "外链菜单"
	menu1.URL = "http://www.qq.com"

	var topCategory models.Category

	topCategory.Name = "职工风采"

	var smallCategory1 models.Category
	smallCategory1.Name = "视频"
	smallCategory1.Higher = &topCategory

	var smallCategory2 models.Category
	smallCategory2.Name = "视频"
	smallCategory2.Higher = &topCategory

	var menu2 models.Menu
	menu2.Category = topCategory

	/*菜单输出*/

	menus := []models.Menu{menu1, menu2}
	var outputMenus []Menu
	for _, menu := range menus {

		var outputMenu Menu

		if menu.Category.ID != 0 {
			category := menu.Category
			var categoryMenu Menu
			categoryMenu.Name = category.Name
			categoryMenu.URL = "/category/" + fmt.Sprintf("%s", category.ID)

			var subCategories []models.Category
			models.DB.Where("higher_id=?", category.ID).Find(&subCategories)
			for _, subCategory := range subCategories {
				var itemMenu Menu
				itemMenu.Name = subCategory.Name
				itemMenu.URL = "/category/" + fmt.Sprintf("%s", subCategory.ID)
				categoryMenu.Menus = append(categoryMenu.Menus, itemMenu)
			}

			outputMenu = categoryMenu

		} else {
			var notCategoryMenu Menu
			notCategoryMenu.URL = menu.URL
			notCategoryMenu.Name = menu.Name
			outputMenu = notCategoryMenu
		}

		outputMenus = append(outputMenus, outputMenu)
	}

	this.ReturnSuccess("menus", outputMenus)

}

func (this *CommonController) Nav() {
	var Menus []*models.Menu

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
}
