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

	models.DB.Select("cover,id,title").Where("category_id=?", 230).Order("created_at desc").Limit(6).Find(&politicalNews)
	models.DB.Select("cover,id,title").Where("category_id=?", 231).Order("created_at desc").Limit(6).Find(&unionNews)
	models.DB.Select("cover,id,title").Where("category_id=?", 233).Order("created_at desc").Limit(6).Find(&grassrootsNews)
	models.DB.Select("cover,id,title").Where("category_id=?", 232).Order("created_at desc").Limit(6).Find(&wechatNews)

	this.Data["politicalNews"] = politicalNews
	this.Data["unionNews"] = unionNews
	this.Data["grassrootsNews"] = grassrootsNews
	this.Data["wechatNews"] = wechatNews
}

func (this *Controller) LoadImageNews() {
	var imageNews []models.Article
	models.DB.Select("cover,id,title").Where("category_id=?", 235).Order("created_at desc").Limit(10).Find(&imageNews)
	this.Data["imageNews"] = imageNews
}

func (this *Controller) LoadPhotoNews() {
	var photoNews []models.Article
	models.DB.Select("cover,id,title").Where("category_id=?", 245).Order("created_at desc").Limit(5).Find(&photoNews)
	this.Data["photoNews"] = photoNews
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

	this.LoadNews()
	this.LoadImageNews()
	this.LoadPhotoNews()

	models.DB.Order("sequence asc").Find(&rotation)
	this.Data["rotations"] = rotation
	this.Data["homes"] = homes

	this.TplName = "web/index.html"
}
