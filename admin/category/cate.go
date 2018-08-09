package category

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/qor"
	"github.com/spf13/cast"
	"github.com/astaxie/beego"
	"strings"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
	"github.com/pkg/errors"
	"fmt"
)

func SetAdmin(adminConfig *admin.Admin) {
	cate := adminConfig.AddResource(&models.Category{}, &admin.Config{Name: "分类管理"})

	cate.SearchAttrs("Name")

	cate.IndexAttrs("ID", "Name", "Sequence", "Category")
	cate.EditAttrs("ID", "Name", "Sequence", "Category", "Higher")
	cate.NewAttrs("ID", "Name", "Sequence", "Category", "Higher")

	cate.Meta(&admin.Meta{Name: "Name",
		Label: "分类名"})

	cate.Meta(&admin.Meta{Name: "Sequence",
		Label: "顺序"})

	//页面分类
	cate.Meta(&admin.Meta{Name: "Category", Label: "类别", Type: "String",
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

	cate.Meta(&admin.Meta{Name: "Higher", Label: "请选择父分类",
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
						return errors.New("请不要选择自身为父分类")
					}

					if cate.ID != 0 {
						var subcate models.Category
						if err := context.GetDB().
							Where("id =?", cate.HigherID).
							First(&subcate).Error; err != nil {
						}

						cate.HigherID = cate.HigherID
						cate.ID = cate.ID
						cate.Higher = &subcate

						context.GetDB().Model(&cate).UpdateColumn("higher_id", cate.HigherID)
						return nil

					}

				}

			}

			return nil
		},
	})

	//重置删除
	cate.Action(
		&admin.Action{
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
			Modes: []string{"show", "menu_item",},
		},
	)
}
