package article

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/media/asset_manager"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	utils2 "github.com/qor/qor/utils"
	"github.com/qor/validations"
	"github.com/spf13/cast"
	"io/ioutil"
	"strings"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/utils"
)

func SetAdmin(adminConfig *admin.Admin) {
	article := adminConfig.AddResource(&models.Article{}, &admin.Config{Name: "文章管理", PageCount: 10})
	//对增删查改的局部显示
	article.IndexAttrs("ID", "Title", "Author", "Cover", "VideoIndex",
		"Editor", "ResponsibleEditor", "Status", "IsIndexUp", "IsIndex", "ReadNum", "Url")

	article.EditAttrs("Title", "Author", "Summary", "Category", "VideoIndex",
		"Cover", "Content", "Editor", "ResponsibleEditor", "Url")

	article.NewAttrs("ID", "Title", "Author", "Summary", "Category", "VideoIndex",
		"Cover", "Content", "Editor", "ResponsibleEditor", "Url")

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
	article.Meta(&admin.Meta{Name: "VideoIndex", Label: "首页封面视频"})
	article.Meta(&admin.Meta{Name: "IsIndexUp", Label: "是否首页置顶"})
	article.Meta(&admin.Meta{Name: "Summary", Label: "文章摘要"})

	article.Meta(&admin.Meta{Name: "Cover", Label: "封面图"})
	article.Meta(&admin.Meta{Name: "Title", Label: "标题"})
	article.Meta(&admin.Meta{Name: "Author", Label: "作者"})
	article.Meta(&admin.Meta{Name: "Editor", Label: "编辑"})
	article.Meta(&admin.Meta{Name: "Source", Label: "来源"})
	article.Meta(&admin.Meta{Name: "ResponsibleEditor", Label: "责任编辑"})
	article.Meta(&admin.Meta{Name: "ReadNum", Label: "阅读数"})
	article.Meta(&admin.Meta{Name: "Url", Label: "转载链接(选填)"})
	article.Meta(&admin.Meta{Name: "IsIndex", Label: "是否显示在主页"})
	//新增的时候的回调
	article.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if a, ok := value.(*models.Article); ok {
				fnameCover := cast.ToString(a.Cover.FileName)
				//调用文件上传函数 更新url
				fmt.Println("this is a debug ------------------------")
				if a.Cover.FileHeader != nil {
					file, err := a.Cover.FileHeader.Open()
					f, err := ioutil.ReadAll(file)

					if err != nil {
						return err
					}
					url, err := utils.UploadFile(fnameCover, f)

					if err != nil {
						return err
					}
					a.Cover.Url = url
				}
				fnameVideo := cast.ToString(a.VideoIndex.FileName)
				if a.VideoIndex.FileHeader != nil {
					file, err := a.VideoIndex.FileHeader.Open()
					f, err := ioutil.ReadAll(file)

					if err != nil {
						return err
					}
					url, err := utils.UploadFile(fnameVideo, f)

					if err != nil {
						return err
					}
					a.VideoIndex.Url = url
				}

				context.GetDB().Where("ID =?", a.CategoryID).First(&a.Category)

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
	//是否显示在首页
	article.Meta(&admin.Meta{Name: "IsIndex", Label: "是否显示在首页", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
		txt := ""
		if v, ok := record.(*models.Article); ok {
			if v.IsIndex == 1 {
				txt = "是"
			} else {
				txt = "否"
			}
		}
		return txt
	}})

	//首页置顶
	article.Meta(&admin.Meta{Name: "IsIndexUp", Label: "是否首页分类置顶", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
		txt := ""
		if v, ok := record.(*models.Article); ok {
			if v.IsIndexUp == 1 {
				txt = "置顶"
			} else {
				txt = "不置顶"
			}
		}
		return txt
	}})
	//添加审核模块
	article.Action(
		&admin.Action{
			Name:  "verify",
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
			Modes: []string{"batch", "show", "menu_item"},
		},
	)
	//添加是否置顶
	article.Action(
		&admin.Action{
			Name:  "isUpIndex",
			Label: "首页置顶/取消置顶",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.Article); ok {
						//执行a.status更新状态
						if a.IsIndexUp == 1 {
							a.IsIndexUp = 0
						} else {
							a.IsIndexUp = 1
						}

						models.DB.Model(&a).Update("IsIndexUp", a.IsIndexUp)

					}
				}
				return nil
			},
			Modes: []string{"show", "menu_item"},
		},
	)
	//重置删除
	article.Action(&admin.Action{
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
		Modes: []string{"show", "menu_item"},
	})
	//是否显示首页
	article.Action(
		&admin.Action{
			Name:  "isIndex",
			Label: "首页显示/首页隐藏",
			Handler: func(argument *admin.ActionArgument) error {
				for _, record := range argument.FindSelectedRecords() {

					if a, ok := record.(*models.Article); ok {
						//执行a.status更新状态
						if a.IsIndex == 1 {
							a.IsIndex = 0
						} else {
							a.IsIndex = 1
						}

						models.DB.Model(&a).Update("IsIndex", a.IsIndex)

					}
				}
				return nil
			},
			Modes: []string{"show", "menu_item"},
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

	//添加字段验证
	article.AddValidator(&resource.Validator{
		Name: "check_article_col",
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {

			if meta := metaValues.Get("Title"); meta != nil {
				if name := utils2.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Title", "标题不能为空")
				}
			}
			if meta := metaValues.Get("Editor"); meta != nil {
				if name := utils2.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Editor", "编辑不能为空")
				}
			}
			if meta := metaValues.Get("ResponsibleEditor"); meta != nil {
				if name := utils2.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "ResponsibleEditor", "责任编辑不能为空")
				}
			}
			if meta := metaValues.Get("Source"); meta != nil {
				if name := utils2.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Title", "来源不能为空")
				}
			}
			return nil
		},
	})
}
