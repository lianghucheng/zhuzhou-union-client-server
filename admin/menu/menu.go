package menu

import (
	"errors"
	"fmt"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego"
)

func SetAdmin(adminConfig *admin.Admin) {
	menu := adminConfig.AddResource(&models.Menu{}, &admin.Config{Name: "导航管理", PageCount: 10})

	//

	menu.SearchAttrs("Name", "ID", "URL", "Category", "Higher", "Sequence")

	menu.IndexAttrs("ID", "Name", "URL", "Category", "Higher", "Sequence")
	menu.EditAttrs("ID", "Name", "URL", "Category", "Higher", "Sequence")
	menu.NewAttrs("ID", "Name", "URL", "Category", "Higher", "Sequence")

	menu.FindManyHandler = func(result interface{}, context *qor.Context) error {
		// find records and decode them to results
		db := context.GetDB()
		if _, ok := db.Get("qor:getting_total_count"); ok {
			return context.GetDB().Count(result).Error
		}
		err := context.GetDB().Set("gorm:order_by_primary_key", "DESC").Find(result).Error
		if menus, ok := result.([]models.Menu); ok {
			beego.Debug(menus[0].Name)
		}

		beego.Debug(result)
		return err
	}

	//重置删除
	menu.Action(&admin.Action{
		Name: "Delete",

		Label: "删除",
		Handler: func(argument *admin.ActionArgument) error {
			for _, record := range argument.FindSelectedRecords() {

				if m, ok := record.(*models.Menu); ok {
					if err := models.DB.Delete(&m).Error; err != nil {
						return err
					}
				}
			}
			return nil
		},
		Modes: []string{"show", "menu_item"},
	})

	menu.Meta(&admin.Meta{Name: "Name",
		Label: "导航名"})
	//URL
	menu.Meta(&admin.Meta{Name: "URL",
		Label: "链接(可不填)"})

	//上级导航
	menu.Meta(&admin.Meta{Name: "Higher",
		Label: "上级导航", Config: &admin.SelectOneConfig{
			Collection: func(_ interface{}, context *admin.Context) (options [][]string) {
				var menus []models.Menu
				context.GetDB().Where("higher_id=?", 0).Find(&menus)
				for _, n := range menus {
					idStr := fmt.Sprintf("%d", n.ID)
					var option = []string{idStr, n.Name}
					options = append(options, option)
				}

				return options
			}, AllowBlank: true, Placeholder: "请选择一个选项"}})

	//栏目
	menu.Meta(&admin.Meta{Name: "Category",
		Label: "对应分类", Config: &admin.SelectOneConfig{
			Collection: func(_ interface{}, context *admin.Context) (options [][]string) {
				var categories []models.Category
				context.GetDB().Where("higher_id=?", 0).Find(&categories)
				for _, n := range categories {
					idStr := fmt.Sprintf("%d", n.ID)
					var option = []string{idStr, n.Name}
					options = append(options, option)
				}

				return options
			}, AllowBlank: true, Placeholder: "请选择一个选项"}})

	//顺序
	menu.Meta(&admin.Meta{Name: "Sequence",
		Label: "菜单顺序"})

	//设置为一级导航
	menu.Action(&admin.Action{
		Name:  "设置导航操作",
		Label: "置为一级导航",
		Handler: func(argument *admin.ActionArgument) error {
			for _, record := range argument.FindSelectedRecords() {

				if m, ok := record.(*models.Menu); ok {

					models.DB.Model(&m).Update("higher_id", 0)
				}
			}
			return nil
		},
		Modes: []string{"batch", "show", "menu_item", "edit"},
	})

	//取消栏目分类关联
	menu.Action(&admin.Action{
		Name:  "取消栏目-分类关联",
		Label: "取消栏目-分类关联",
		Handler: func(argument *admin.ActionArgument) error {
			for _, record := range argument.FindSelectedRecords() {

				if c, ok := record.(*models.Menu); ok {

					models.DB.Model(&c).Update("category_id", 0)
				}
			}
			return nil
		},
		Modes: []string{"batch", "show", "menu_item", "edit"},
	})

	//添加分类时父分类不能为自己
	menu.AddProcessor(&resource.Processor{
		Name: "process_menu_data",
		Handler: func(record interface{}, values *resource.MetaValues, context *qor.Context) error {
			if m, ok := record.(*models.Menu); ok {
				if m.Higher != nil {
					if m.ID == m.Higher.ID {
						return errors.New("请不要选择自身为上级导航")
					}
					if err := context.GetDB().
						Where("id =?", m.HigherID).
						First(&m.Higher).Error; err != nil {
					}
					return nil
				}
			}
			return nil
		},
	})

	//字段验证
	menu.AddValidator(&resource.Validator{
		Name: "check_menu_col",
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			//if meta := metaValues.Get("Name"); meta != nil {
			//	if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
			//		return validations.NewError(record, "Name", "导航名不能为空")
			//	}
			//}
			//
			//url := metaValues.Get("Url");bee run
			//if utils.ToString(url.Value) == "" && len(cate) == 0 {
			//	return errors.New("链接和分类必须任选其一")
			//}

			return nil
		},
	})
}
