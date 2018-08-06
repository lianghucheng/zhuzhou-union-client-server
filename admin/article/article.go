package article

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/media/asset_manager"
	"fmt"
)

func SetAdmin(adminConfig *admin.Admin) {
	article := adminConfig.AddResource(&models.Article{})

	article.IndexAttrs("ID", "Title", "Author", "Cover", "Content", "Editor", "ResponsibleEditor", "Category.Name")

	//富文本
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

	//下拉选项文章作者
	article.Meta(&admin.Meta{Name: "User", Label: "User", Type: "select_many",
		Config: &admin.SelectOneConfig{
			Collection: func(_ interface{}, context *admin.Context) (options [][]string) {
				var users []models.User
				context.GetDB().Find(&users)

				for _, n := range users {
					idStr := fmt.Sprintf("%d", n.ID)
					var option = []string{idStr, n.Username}
					options = append(options, option)
				}

				return options
			},
		},
	})

	//顶部扩展栏 通过选择选项来执行
	article.Action(
		&admin.Action{
			Name:  "enable",
			Label: "审核通过/撤销通过",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(models.Article); ok {
						//执行a.status更新状态
						fmt.Println(a)
					}

					fmt.Println(record)
				}
				return nil
			},
			Modes: []string{"batch", "edit", "show"},
		},
	)

}
