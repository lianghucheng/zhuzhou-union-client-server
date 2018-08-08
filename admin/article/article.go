package article

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/media/asset_manager"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"fmt"
	"github.com/jinzhu/gorm"
	"zhuzhou-union-client-server/utils"
	"github.com/spf13/cast"
	"io/ioutil"
)

func SetAdmin(adminConfig *admin.Admin) {
	article := adminConfig.AddResource(&models.Article{}, &admin.Config{Name: "文章管理"})
	//对增删查改的局部显示
	article.IndexAttrs("ID", "Title", "Author", "Cover", "Editor", "ResponsibleEditor", "Status")
	article.EditAttrs("Title", "Author", "Category", "Cover", "Content", "Editor", "ResponsibleEditor")
	article.NewAttrs("ID", "Title", "Author", "Category", "Cover", "Content", "Editor", "ResponsibleEditor")

	//添加富文本
	assetManager := adminConfig.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	article.Meta(&admin.Meta{Name: "Content", Label: "内容", Config: &admin.RichEditorConfig{
		AssetManager: assetManager,
		Plugins: []admin.RedactorPlugin{
			{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
			{Name: "table", Source: "/admin/assets/javascripts/qor_kindeditor.js"},
		},
		Settings: map[string]interface{}{
			"medialibraryUrl": "/admin/product_images",
		},
	}})
	article.Meta(&admin.Meta{Name: "Content", Label: "内容", Type: "kindeditor"})

	article.Meta(&admin.Meta{Name: "Cover", Label: "封面图"})
	article.Meta(&admin.Meta{Name: "Title", Label: "标题"})
	article.Meta(&admin.Meta{Name: "Author", Label: "作者"})
	article.Meta(&admin.Meta{Name: "Editor", Label: "编辑"})
	article.Meta(&admin.Meta{Name: "Source", Label: "来源"})
	article.Meta(&admin.Meta{Name: "ResponsibleEditor", Label: "责任编辑"})

	//新增的时候的回调
	article.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if a, ok := value.(*models.Article); ok {
				fmt.Println("thiss is a :", a)
				fname := cast.ToString(a.Cover.FileName)
				//调用文件上传函数 更新url
				if a.Cover.FileHeader != nil {
					file, err := a.Cover.FileHeader.Open()
					f, err := ioutil.ReadAll(file)

					if err != nil {
						return err
					}
					url, err := utils.UploadFile(fname, f)

					if err != nil {
						return err
					}
					a.Cover.Url = url
				}

				if a.Category != nil {
					a.CategoryID = a.Category.ID
				}

			}
			return nil
		},
	})

	//重置Status显示
	article.Meta(&admin.Meta{Name: "Status", Label: "审核状态", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
		txt := ""
		if v, ok := record.(*models.Article); ok {
			if v.Status == 1 {

				txt = "已审核"
			} else {
				txt = "未审核"
			}
		}
		return txt
	}})

	//添加审核模块
	article.Action(
		&admin.Action{
			Name:  "enable",
			Label: "审核/撤销",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.Article); ok {
						//执行a.status更新状态
						if a.Status == 1 {
							a.Status = 0
						} else {
							a.Status = 1
						}
						models.DB.Model(&a).Update("status", a.Status)

					}
				}
				return nil
			},
			Modes: []string{"batch", "show", "menu_item", "edit"},
		},
	)

	//重置删除
	article.Action(
		&admin.Action{
			Name:  "Delete",
			Label: "删除",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.Article); ok {
						if err := models.DB.Delete(&a).Error; err != nil {
							return err
						}
					}
				}
				return nil
			},
			Modes: []string{"show", "menu_item",},
		},
	)

	//添加搜索

	article.SearchAttrs("Title", "Content", "Editor", "ResponsibleEditor", "Author", "ID")

	//添加过滤条件
	article.Scope(&admin.Scope{Name: "已审核", Group: "审核状态", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("status = ?", "1")
	}})

	article.Scope(&admin.Scope{Name: "未审核", Group: "审核状态", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("status = ?", "0")
	}})

	//添加分类选项
	article.Meta(&admin.Meta{Name: "Category", Label: "请选择分类"})
}
