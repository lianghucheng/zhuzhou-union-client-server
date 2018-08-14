package home

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor"
	"github.com/spf13/cast"
	"io/ioutil"
	"zhuzhou-union-client-server/utils"
	"github.com/qor/validations"
	utils2 "github.com/qor/qor/utils"
)

func SetAdmin(adminConfig *admin.Admin) {
	home := adminConfig.AddResource(&models.Home{}, &admin.Config{Menu: []string{"首页管理"}, Name: "中间文章模块管理"})

	//对增删查改的局部显示
	home.IndexAttrs("ID", "Category", "IndexArticle", "Position", "Layout")
	home.EditAttrs("Category", "IndexArticle", "Position", "Layout")
	home.NewAttrs("Category", "IndexArticle", "Position", "Layout")

	home.Meta(&admin.Meta{Name: "Category", Label: "首页分类"})
	home.Meta(&admin.Meta{Name: "Position", Label: "具体位置"})
	home.Meta(&admin.Meta{Name: "Layout", Label: "模块位置"})
	home.Meta(&admin.Meta{Name: "IndexArticle", Label: "单个分类置顶文章"})

	imageLinks := adminConfig.AddResource(&models.ImageLinks{}, &admin.Config{Menu: []string{"首页管理"}, Name: "底部图片链接管理"})
	imageLinks.IndexAttrs("ID", "Url", "Image")
	imageLinks.EditAttrs("Url", "Image")
	imageLinks.NewAttrs("Url", "Image")

	imageLinks.Meta(&admin.Meta{Name: "Url", Label: "具体链接地址"})
	imageLinks.Meta(&admin.Meta{Name: "Image", Label: "显示图片"})

	imageLinks.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if i, ok := value.(*models.ImageLinks); ok {
				fname := cast.ToString(i.Image.FileName)
				//调用文件上传函数 更新url
				if i.Image.FileHeader != nil {
					file, err := i.Image.FileHeader.Open()
					f, err := ioutil.ReadAll(file)

					if err != nil {
						return err
					}
					url, err := utils.UploadFile(fname, f)

					if err != nil {
						return err
					}
					i.Image.Url = url
				}

			}
			return nil
		},
	})

	boxLinks := adminConfig.AddResource(&models.BoxLinks{}, &admin.Config{Menu: []string{"首页管理"}, Name: "底部下拉链接管理"})
	boxLinks.IndexAttrs("ID", "Name", "Url", "Position")
	boxLinks.EditAttrs("Name", "Url", "Position")
	boxLinks.NewAttrs("Name", "Url", "Position")

	boxLinks.Meta(&admin.Meta{Name: "Name", Label: "名称"})
	boxLinks.Meta(&admin.Meta{Name: "Url", Label: "链接"})
	boxLinks.Meta(&admin.Meta{Name: "Position", Label: "下拉框位置"})

	imageLinks.AddValidator(&resource.Validator{
		Name: "check_article_col",
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {

			if meta := metaValues.Get("position"); meta != nil {

				if position := utils2.ToInt(meta.Value); position < 0 {
					return validations.NewError(record, "Position", "请输入大于0小于显示条数的数量")
				}

			}
			return nil
		},
	})

	boxLinks.AddValidator(&resource.Validator{
		Name: "check_article_col",
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {

			if meta := metaValues.Get("position"); meta != nil {

				if position := utils2.ToInt(meta.Value); position <= 0 || position > 3 {
					return validations.NewError(record, "Position", "请输入1-3的位置")
				}

			}
			return nil
		},
	})

}
