package main

import (
	"encoding/csv"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"zhuzhou-union-client-server/models"
)

type Class struct {
	ID   string
	Name string
}

type Article struct {
	ID             string
	Tid            string
	Title          string
	Intro          string
	ArticleContent string
	Author         string
	Origin         string
	AddDate        string
	ReadTimes      string
	PhotoUrl       string
	Inputer        string
}

func transfer() {
	cntb1, err := ioutil.ReadFile("KS_Class.csv")
	if err != nil {
		panic(err)
	}
	r1 := csv.NewReader(strings.NewReader(string(cntb1)))
	classCSV, _ := r1.ReadAll()

	cntb2, err := ioutil.ReadFile("KS_Article.csv")
	if err != nil {
		panic(err)
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb2)))
	articleCSV, _ := r2.ReadAll()

	classes := make(map[string]string)
	categories := make(map[string]uint)

	var articles []Article
	for index, class := range classCSV {
		if index == 0 {
			continue
		}
		classes[class[0]] = class[3]
	}

	for index, article := range articleCSV {
		if index == 0 {
			continue
		}
		var a Article
		a.ID = article[0]
		a.Tid = article[1]
		a.Title = article[4]
		a.Intro = article[6]
		a.ArticleContent = article[10]
		a.Author = article[11]
		a.Origin = article[12]
		a.AddDate = article[15]
		a.ReadTimes = article[36]
		a.PhotoUrl = article[42]
		a.Inputer = article[43]
		articles = append(articles, a)
	}

	var cates []models.Category
	models.DB.Where("higher_id <> ?", 0).Find(&categories)
	for _, category := range cates {
		categories[category.Name] = category.ID
	}

	beego.Debug("共计 ", len(articles), " 篇文章，开始写入")
	i := 0
	for _, article := range articles {
		if name, ok := classes[article.Tid]; ok {
			if id, ok := categories[name]; ok {
				var insert models.Article
				insert.Title = article.Title
				insert.CategoryID = id
				insert.Author = article.Author
				insert.Source = article.Origin
				insert.Summary = article.Intro
				insert.Content = article.ArticleContent
				insert.CreatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", article.AddDate, time.Local)
				if article.PhotoUrl != "" {
					insert.Cover.Url = article.PhotoUrl
				}
				insert.Editor = article.Inputer
				ReadNum, _ := strconv.ParseInt(article.ReadTimes, 10, 64)
				insert.ReadNum = uint(ReadNum)
				id, _ := strconv.ParseInt(article.ID, 10, 64)
				insert.ID = uint(id)
				models.DB.Save(&insert)
				i++
				beego.Debug("已写入", i, "篇文章")
			}
		}
	}

}
