package controllers

import (
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego/utils/pagination"
	"time"
)

type SearchController struct {
	Common
}

//@router /search [*]
func (this *SearchController) Search() {

	str := this.GetString("str")

	startTime := this.GetString("startTime")
	endTime := this.GetString("endTime")

	pers := 6

	qs := models.DB.Select("id,cover,summary,title,author,created_at").
		Model(models.Article{})
	if str != "" {
		qs = qs.Where("title like ? or content like ? or author like ? or editor like ?",
			"%"+str+"%", "%"+str+"%", "%"+str+"%", "%"+str+"%")
	}
	if startTime != "" && endTime != "" {
		sT, err := time.Parse("2006-01-02 15:04:05", startTime)
		eT, err := time.Parse("2006-01-02 15:04:05", endTime)
		if err == nil {
			qs.Where("created_at between ? and ?", sT, eT)
		}
	}

	cnt := 0
	qs.Count(&cnt)

	pager := pagination.NewPaginator(this.Ctx.Request, pers, cnt)
	qs = qs.Order("created_at desc").Limit(pers).Offset(pager.Offset())
	var articles []*models.Article
	qs.Find(&articles)

	this.Data["articles"] = articles
	this.Data["paginator"] = pager

	this.TplName = "search.html"
}
