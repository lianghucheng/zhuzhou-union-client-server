package home

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
)

func (this *Controller) LoadNews() {

	var politicalNews []models.Article
	var unionNews []models.Article
	var grassrootsNews []models.Article
	var wechatNews []models.Article

	models.DB.Where("category_id=?", 230).Order("created_at desc").Limit(6).Find(&politicalNews)
	models.DB.Where("category_id=?", 231).Order("created_at desc").Limit(6).Find(&unionNews)
	models.DB.Where("category_id=?", 233).Order("created_at desc").Limit(6).Find(&grassrootsNews)
	models.DB.Where("category_id=?", 232).Order("created_at desc").Limit(6).Find(&wechatNews)

	this.Data["politicalNews"] = politicalNews
	this.Data["unionNews"] = unionNews
	this.Data["grassrootsNews"] = grassrootsNews
	this.Data["wechatNews"] = wechatNews
}

//@router /	[*]
func (this *Controller) Index() {

	var homes []*models.Home
	var rotation []*models.Rotation
	var indexPer int

	indexPer, _ = beego.AppConfig.Int("indexPer")
	if indexPer == 0 {
		indexPer = 5
	}

	models.DB.Order("sequence asc").Find(&rotation)
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

	this.LoadNews()

	this.Data["rotations"] = rotation
	this.Data["imageLinks"] = imageLinks
	this.Data["homes"] = homes
	this.Data["boxLinks"] = boxLinks
	this.TplName = "web/index.html"
}
