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
	/*home := adminConfig.AddResource(&models.Home{}, &admin.Config{Menu: []string{"首页管理"}, Name: "中间文章模块管理", PageCount: 10})

	//对增删查改的局部显示
	home.IndexAttrs("ID", "Name", "Category", "Position", "Layout", "Url")
	home.EditAttrs("Name", "Category", "Position", "Layout", "Url")
	home.NewAttrs("Name", "Category", "Position", "Layout", "Url")

	home.Meta(&admin.Meta{Name: "Name", Label: "分类名"})
	home.Meta(&admin.Meta{Name: "Url", Label: "具体链接(可不填)"})
	home.Meta(&admin.Meta{Name: "Category", Label: "首页分类"})
	home.Meta(&admin.Meta{Name: "Position", Label: "具体位置"})
	home.Meta(&admin.Meta{Name: "Layout", Label: "模块位置"})*/

	rotation := adminConfig.AddResource(&models.Rotation{}, &admin.Config{Menu: []string{"首页管理"}, Name: "轮播图管理", PageCount: 10})

	rotation.IndexAttrs("ID", "Url", "Position", "Sequence")
	rotation.EditAttrs("Url", "Position", "Sequence")
	rotation.NewAttrs("Url", "Position", "Sequence")

	rotation.Meta(&admin.Meta{Name: "Url", Label: "轮播图"})
	rotation.Meta(&admin.Meta{Name: "Position", Label: "位置"})
	rotation.Meta(&admin.Meta{Name: "Sequence", Label: "顺序"})

	rotation.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if r, ok := value.(*models.Rotation); ok {
				fname := cast.ToString(r.Url.FileName)
				//调用文件上传函数 更新url
				if r.Url.FileHeader != nil {
					file, err := r.Url.FileHeader.Open()
					f, err := ioutil.ReadAll(file)

					if err != nil {
						return err
					}
					url, err := utils.UploadFile(fname, f)

					if err != nil {
						return err
					}
					r.Url.Url = url
				}
			}
			return nil
		},
	})

	imageLinks := adminConfig.AddResource(&models.ImageLinks{}, &admin.Config{Menu: []string{"首页管理"}, Name: "底部图片链接管理", PageCount: 10})
	imageLinks.IndexAttrs("ID", "Url", "Image", "Position")
	imageLinks.EditAttrs("Url", "Image", "Position")
	imageLinks.NewAttrs("Url", "Image", "Position")

	imageLinks.Meta(&admin.Meta{Name: "Url", Label: "具体链接地址"})
	imageLinks.Meta(&admin.Meta{Name: "Image", Label: "显示图片"})
	imageLinks.Meta(&admin.Meta{Name: "Position", Label: "位置"})

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

	boxLinks := adminConfig.AddResource(&models.BoxLinks{}, &admin.Config{Menu: []string{"首页管理"}, Name: "底部下拉链接管理", PageCount: 10})
	boxLinks.IndexAttrs("ID", "Name", "Url", "Position", "IsUp")
	boxLinks.EditAttrs("Name", "Url", "Position", "IsUp")
	boxLinks.NewAttrs("Name", "Url", "Position", "IsUp")

	boxLinks.Meta(&admin.Meta{Name: "Name", Label: "名称"})
	boxLinks.Meta(&admin.Meta{Name: "Url", Label: "链接"})
	boxLinks.Meta(&admin.Meta{Name: "Position", Label: "下拉框位置"})
	boxLinks.Meta(&admin.Meta{Name: "IsUp", Label: "是否最先显示"})

	//重置IsUp显示
	boxLinks.Meta(&admin.Meta{Name: "IsUp", Label: "是否最先显示", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
		txt := ""
		if v, ok := record.(*models.BoxLinks); ok {
			if v.IsUp == 1 {

				txt = "显示"
			} else {
				txt = "不显示"
			}
		}
		return txt
	}})

	//添加是否置顶显示
	boxLinks.Action(
		&admin.Action{
			Name:  "isUpb",
			Label: "显示/取消",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.BoxLinks); ok {
						//执行a.status更新状态
						if a.IsUp == 1 {
							a.IsUp = 0
						} else {
							a.IsUp = 1
						}
						models.DB.Model(&a).Update("IsUp", a.IsUp)

					}
				}
				return nil
			},
			Modes: []string{"show", "menu_item"},
		},
	)

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

	qrCode := adminConfig.AddResource(&models.QrCode{}, &admin.Config{Menu: []string{"首页管理"}, Name: "首页二维码管理"})
	qrCode.IndexAttrs("ID", "CodeImage")
	qrCode.EditAttrs("CodeImage")
	qrCode.NewAttrs("CodeImage")
	qrCode.Meta(&admin.Meta{Name: "CodeImage", Label: "二维码图片"})

	qrCode.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if a, ok := value.(*models.QrCode); ok {
				fnameCover := cast.ToString(a.CodeImage.FileName)
				//调用文件上传函数 更新url
				if a.CodeImage.FileHeader != nil {
					file, err := a.CodeImage.FileHeader.Open()
					f, err := ioutil.ReadAll(file)

					if err != nil {
						return err
					}
					url, err := utils.UploadFile(fnameCover, f)

					if err != nil {
						return err
					}
					a.CodeImage.Url = url
				}

			}
			return nil
		},
	})

}
