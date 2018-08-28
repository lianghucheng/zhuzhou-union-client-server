package category

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
	"github.com/spf13/cast"
	"strings"
	"zhuzhou-union-client-server/models"
)

func SetAdmin(adminConfig *admin.Admin) {
	cate := adminConfig.AddResource(&models.Category{}, &admin.Config{Name: "分类管理", PageCount: 10})

	cate.SearchAttrs("Name", "Category", "Higher", "ID")

	cate.IndexAttrs("ID", "Name", "Sequence", "Category", "Higher", "IsSubmission")
	cate.EditAttrs("ID", "Name", "Sequence", "Category", "Higher")
	cate.NewAttrs("ID", "Name", "Sequence", "Category", "Higher")

	//分类名
	cate.Meta(&admin.Meta{Name: "Name",
		Label: "分类名"})
	//顺序
	cate.Meta(&admin.Meta{Name: "Sequence",
		Label: "顺序"})

	//上级分类
	cate.Meta(&admin.Meta{Name: "Higher",
		Label: "上级分类", Config: &admin.SelectOneConfig{
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

	cate.Meta(&admin.Meta{Name: "Special",
		Label: "是否在文章侧边栏显示"})

	//页面分类
	cate.Meta(&admin.Meta{Name: "Category",
		Label: "类别",
		Type: "String",
		FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
			txt := ""
			cateNames := strings.Split(beego.AppConfig.String("catename"), ",")
			if v, ok := record.(*models.Category); ok {
				for index, _ := range cateNames {
					if v.Category == index {
						txt = cateNames[index]
					}
				}
				return txt
			}

			return nil
		}})

	//添加分类选项
	cate.Meta(&admin.Meta{Name: "Category", Label: "请选择类别", Type: "select_many",
		Config: &admin.SelectOneConfig{
			Collection: func(_ interface{}, context *admin.Context) (options [][]string) {
				var cateName []string

				catenames := beego.AppConfig.String("catename")
				cateName = strings.Split(catenames, ",")
				for index, _ := range cateName {
					var option = []string{cast.ToString(index), cateName[index]}
					options = append(options, option)
				}
				return options
			},
		},
	})

	//新增时候的验证
	cate.AddValidator(&resource.Validator{
		Name: "check_has_name",
		Handler: func(record interface{}, values *resource.MetaValues, context *qor.Context) error {
			if meta := values.Get("Name"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Name", "分类名不能为空")
				}
			}
			if meta := values.Get("Sequence"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Sequence", "顺序不能为空")
				}
			}
			return nil
		},
	})

	//添加分类时父分类不能为自己
	cate.AddProcessor(&resource.Processor{
		Name: "process_cate_data",

		Handler: func(record interface{}, values *resource.MetaValues, context *qor.Context) error {

			if cate, ok := record.(*models.Category); ok {
				if cate.Higher != nil {
					fmt.Println("not nil higher ", cate)
					if cate.ID == cate.Higher.ID {
						return errors.New("请不要选择自身为上级分类")
					}

					//var subcate models.Category

					fmt.Println("start :", cate.ID, cate.HigherID, cate.Higher)
					if err := context.GetDB().
						Where("id =?", cate.HigherID).
						First(&cate.Higher).Error; err != nil {
					}

					fmt.Println("end :", cate.ID, cate.HigherID, cate.Higher)
					return nil

				}

			}

			return nil
		},
	})

	//重置删除
	cate.Action(&admin.Action{
		Name:  "Delete",
		Label: "删除",
		Handler: func(argument *admin.ActionArgument) error {
			for _, record := range argument.FindSelectedRecords() {

				if a, ok := record.(*models.Category); ok {
					if err := models.DB.Delete(&a).Error; err != nil {
						return err
					}
				}
			}
			return nil
		},
		Modes: []string{"show", "menu_item"},
	})

	//将节点设置为根分类
	cate.Action(&admin.Action{
		Name:  "设置一级分类操作",
		Label: "置为一级分类",
		Handler: func(argument *admin.ActionArgument) error {
			for _, record := range argument.FindSelectedRecords() {

				if c, ok := record.(*models.Category); ok {

					models.DB.Model(&c).Update("higher_id", 0)
				}
			}
			return nil
		},
		Modes: []string{"batch", "show", "menu_item", "edit"},
	})

	//是否设置为侧边栏分类
	cate.Action(
		&admin.Action{
			Name:  "设置侧边栏操作",
			Label: "置为侧边栏/取消",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.Category); ok {
						//执行a.status更新状态
						if a.Special == 1 {
							a.Special = 0
						} else {
							a.Special = 1
						}
						models.DB.Model(&a).Update("Special", a.Special)

					}
				}
				return nil
			},
			Modes: []string{"show", "menu_item", "edit"},
		},
	)

	cate.Meta(&admin.Meta{Name: "IsSubmission", Label: "是否为投稿分类", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
		txt := ""
		if v, ok := record.(*models.Category); ok {
			if v.IsSubmission == 1 {

				txt = "投稿分类"
			} else {
				txt = "非投稿分类"
			}
		}
		return txt
	}})
	cate.Action(
		&admin.Action{
			Name:  "置为投稿分类",
			Label: "置为投稿分类/取消",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.Category); ok {
						//执行a.status更新状态
						if a.IsSubmission == 1 {
							a.IsSubmission = 0
						} else {
							a.IsSubmission = 1
						}
						models.DB.Model(&a).Update("is_submission", a.IsSubmission)

					}
				}
				return nil
			},
			Modes: []string{"batch", "show", "menu_item"},
		},
	)
}
