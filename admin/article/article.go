package article

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/media/asset_manager"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"fmt"
)

func SetAdmin(adminConfig *admin.Admin) {
	article := adminConfig.AddResource(&models.Article{})
	//对增删查改的局部显示
	article.IndexAttrs("ID", "Title", "Author", "Cover", "Content", "Editor", "ResponsibleEditor", "Status")
	article.EditAttrs("Title", "Author", "Cover", "Content", "Editor", "ResponsibleEditor")
	article.NewAttrs("ID", "Title", "Author", "Cover", "Content", "Editor", "ResponsibleEditor")

	//添加富文本
	assetManager := adminConfig.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	article.Meta(&admin.Meta{Name: "Content", Config: &admin.RichEditorConfig{
		AssetManager: assetManager,
		Plugins: []admin.RedactorPlugin{
			{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
			{Name: "table", Source: "/admin/assets/javascripts/qor_kindeditor.js"},
		},
		Settings: map[string]interface{}{
			"medialibraryUrl": "/admin/product_images",
		},
	}})
	article.Meta(&admin.Meta{Name: "Content", Type: "kindeditor"})

	article.Meta(&admin.Meta{Name: "Cover"})


	//新增的时候的回调
	article.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			fmt.Println("--------------------")
			if a, ok := value.(*models.Article); ok {

				//调用文件上传函数 更新url
				fmt.Println(a.Cover.FileHeader)
			}
			return nil
		},
	})

	//重置Status显示
	article.Meta(&admin.Meta{Name: "Status", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
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


}
