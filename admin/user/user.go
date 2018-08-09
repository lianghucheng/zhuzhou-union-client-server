package user

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor"
	"zhuzhou-union-client-server/utils"
	"github.com/jinzhu/gorm"
)

func SetAdmin(adminConfig *admin.Admin) {
	user := adminConfig.AddResource(&models.User{},&admin.Config{Name:"用户管理"})

	user.IndexAttrs("Username","Password","Prioty")
	user.SearchAttrs("Username")
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
