package user

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/astaxie/beego"
	"github.com/qor/qor"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/resource"
	"zhuzhou-union-client-server/utils"
)

func SetAdmina(adminConfig *admin.Admin) {
	user := adminConfig.AddResource(&models.User{},&admin.Config{Name:"用户管理"})
	beego.Debug(user)

	user.IndexAttrs("Username","Password","Prioty")
	user.Meta(&admin.Meta{Name:"Username",Label:"用户名"})
	user.Meta(&admin.Meta{Name:"Password",Label:"密码"})
	user.Meta(&admin.Meta{Name:"Prioty",Label:"权限"})
	user.AddProcessor(&resource.Processor{
		Name: "process_user_data",
		Handler: func(val interface{}, values *resource.MetaValues, context *qor.Context) error {
			if user, ok := val.(*models.User); ok {
				user.Password = utils.Md5(user.Password) // do something...
			}
			return nil
		},
	})

	user.Scope(&admin.Scope{Name: "权限一", Group: "权限等级", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("prioty = ?", 1)
	}})

	user.Scope(&admin.Scope{Name: "权限二", Group: "权限等级", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("prioty = ?", 2)
	}})
	user.Scope(&admin.Scope{Name: "权限三", Group: "权限等级", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("prioty = ?", 3)
	}})
}
/*
1。显示每个用户提交的文章的数量（资源）
2。点击出现符合条件的文章列表（资源）
3。再点击可以编辑文章
 */