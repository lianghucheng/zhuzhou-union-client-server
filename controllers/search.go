package controllers

import (
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego/utils/pagination"
)

type SearchController struct {
	beego.Controller
}


//@router /search [*]
func (this *SearchController) Search() {

	str:=this.GetString("str")

	pers := 6
	qs := models.DB.Select("id,cover,summary,title,author,created_at").
		Model(models.Article{}).
		Where("title like ? or content like ? or author like ? or editor like ?",
			str,str,str,str)
	cnt := 0
	qs.Count(&cnt)

	pager := pagination.NewPaginator(this.Ctx.Request, pers, cnt)
	qs = qs.Order("created_at desc").Limit(pers).Offset(pager.Offset())
	var articles []*models.Article
	qs.Find(&articles)

	this.Data["articles"] = articles
	this.Data["paginator"] = pager



	this.TplName="search.html"
}