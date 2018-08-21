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
	"reflect"
)

func SetAdmin(adminConfig *admin.Admin) {
	adminConfig.DB=adminConfig.DB.Where("name <> 根")
	cate := adminConfig.AddResource(&models.Category{}, &admin.Config{Name: "分类管理", PageCount: 10})

	cate.SearchAttrs("Name", "Category", "Higher", "ID")

	cate.IndexAttrs("ID", "Name", "Sequence", "Category", "Higher", "Special")
	cate.EditAttrs("ID", "Name", "Sequence", "Category", "Higher", "Special")
	cate.NewAttrs("ID", "Name", "Sequence", "Category", "Higher", "Special")

	//分类名
	cate.Meta(&admin.Meta{Name: "Name",
		Label: "分类名"})
	//顺序
	cate.Meta(&admin.Meta{Name: "Sequence",
		Label: "顺序"})

	//上级分类
	cate.Meta(
		&admin.Meta{
			Name:  "Higher",
			Label: "上级栏目",
			FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
				if a, ok := record.(*models.Category); ok {
					if a.HigherID == 0 {
						return "此栏目为一级分类"
					}else{
						cate:=models.Category{}
						if err:=models.DB.Where("id = ?",a.HigherID).Find(&cate).Error;err!=nil{
							beego.Error("栏目过滤异常",err)
							return
						}
						return cate.Name
					}
				}
				return
			},
		},
	)

	cate.Meta(
		&admin.Meta{
			Name:  "Special",
			Label: "是否在文章侧边栏显示",
			Setter: func(val interface{}, values *resource.MetaValue, context *qor.Context) {
				if category, ok := val.(*models.Category); ok {
					beego.Debug("--------------")
					beego.Debug(val)
					beego.Debug(values)
					if values.Name == beego.AppConfig.String("userPowerField") {
						if a, ok := values.Value.([]string); ok {
							beego.Debug(a[0])
							if a[0] == "是" {
								var special []int
								special = append(special, 1)
								beego.Debug(special)
								category.Special = 1
								values.Value = special
							}
							if a[0] == "否" {
								var special []int
								special = append(special, 2)
								beego.Debug(special)
								category.Special = 2
								values.Value = special
							}
						}
					} else {
						beego.Debug(reflect.TypeOf(values.Value))
						beego.Debug("性别断言：", ok)
						if ok {
							a := values.Value.([]string)
							beego.Debug(a[0])
							if a[0] == "是" {
								var special []int
								special = append(special, 1)
								beego.Debug(special)
								category.Special = 1
								values.Value = special
							}
							if a[0] == "否" {
								var special []int
								special = append(special, 2)
								beego.Debug(special)
								category.Special = 0
								values.Value = special
							}
						}
					}
				}
				return
			},
			Type: "text",
			Config: &admin.SelectOneConfig{
				Collection: []string{
					"是",
					"否",
				},
			},
			FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
				if a, ok := record.(*models.Category); ok {
					if a.Special == 0 {
						return "否"
					}
					if a.Special == 1 {
						return "是"
					}
				}
				return beego.AppConfig.String("paseAdminERR")
			},
		},

	)
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
		Modes: []string{"show", "menu_item",},
	}, )

	//将节点设置为根分类
	cate.Action(
		&admin.Action{
			Name:  "enable",
			Label: "置为一级分类",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if c, ok := record.(*models.Category); ok {

						models.DB.Model(&c).Update("higher_id", 0)
					}
				}
				return nil
			},
			Modes: []string{"batch", "show", "menu_item"},
		},
	)

	//是否设置为侧边栏分类
	cate.Action(
		&admin.Action{
			Name:  "isSpecialCate",
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
			Modes: []string{"show", "menu_item"},
		},
	)
}
