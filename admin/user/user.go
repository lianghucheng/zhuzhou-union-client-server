package user

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"zhuzhou-union-client-server/utils"
	"github.com/qor/roles"
)

func SetAdmin(adminConfig *admin.Admin) {

	user := adminConfig.AddResource(
		&models.User{},
		&admin.Config{
			Name: "用户管理",
			Permission: roles.
				Allow(roles.Update, "admin").
				Allow(roles.Delete, "admin").
				Allow(roles.Create, "admin").
				Allow(roles.Read, roles.Anyone),
		},
	)

	user.IndexAttrs("ID", "NickName", "UserName", "Sex", "Prioty")
	user.EditAttrs("-ID", "-Password", "-Icon")
	user.SearchAttrs("Username")

	user.Meta(&admin.Meta{Name: "Username", Label: "用户名"})
	user.Meta(&admin.Meta{Name: "NickName", Label: "昵称"})
	user.Meta(&admin.Meta{Name: "Sex", Label: "性别", Type: "String", FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
		txt := ""
		if v, ok := record.(*models.User); ok {
			if v.Sex == 1 {
				txt = "男"
			} else if v.Sex == 2 {
				txt = "女"
			} else {
				txt = "未知"
			}
		}
		return txt
	}})
	user.Meta(&admin.Meta{Name: "Prioty", Label: "权限"})

	user.Meta(
		&admin.Meta{
			Name:  "Prioty",
			Type:  "text",
			Label: "身份",
			FormattedValuer: func(record interface{}, context *qor.Context) (result interface{}) {
				if a, ok := record.(*models.User); ok {
					if a.Prioty == 1 {
						return beego.AppConfig.String("rootAdmin")
					}
					if a.Prioty == 2 {
						return beego.AppConfig.String("commonAdmin")
					}
					if a.Prioty == 3 {
						return beego.AppConfig.String("user")
					}
				}
				return beego.AppConfig.String("paseAdminERR")
			},
			Config: &admin.SelectOneConfig{
				Collection: []string{
					beego.AppConfig.String("rootAdmin"),
					beego.AppConfig.String("commonAdmin"),
					beego.AppConfig.String("user"),
				},
			},
			Setter: func(val interface{}, values *resource.MetaValue, context *qor.Context) {
				if user, ok := val.(*models.User); ok {
					beego.Debug("--------------")
					beego.Debug(val)
					beego.Debug(values)
					if values.Name == beego.AppConfig.String("userPowerField") {
						if a, ok := values.Value.([]string); ok {
							beego.Debug(a[0])
							if a[0] == beego.AppConfig.String("rootAdmin") {
								var prioty []int
								prioty = append(prioty, 1)
								beego.Debug(prioty)
								user.Prioty = 1
								values.Value = prioty
							}
							if a[0] == beego.AppConfig.String("commonAdmin") {
								var prioty []int
								prioty = append(prioty, 2)
								beego.Debug(prioty)
								user.Prioty = 2
								values.Value = prioty
							}
							if a[0] == beego.AppConfig.String("user") {
								var prioty []int
								prioty = append(prioty, 3)
								beego.Debug(prioty)
								user.Prioty = 3
								values.Value = prioty
							}
						}
					}
				}
				return
			},
		})


	user.AddProcessor(&resource.Processor{
		Name: "process_user_data",
		Handler: func(val interface{}, values *resource.MetaValues, context *qor.Context) error {
			if user, ok := val.(*models.User); ok {
				beego.Debug(context)
				beego.Debug("--------------")
				beego.Debug(user.Password)
				user.Password = utils.Md5(user.Password) // do something...
				beego.Debug(user.Password)
			}
			return nil
		},
	})

	user.Scope(&admin.Scope{Name: beego.AppConfig.String("rootAdmin"), Group: "权限等级", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("prioty = ?", 1)
	}})

	user.Scope(&admin.Scope{Name: beego.AppConfig.String("commonAdmin"), Group: "权限等级", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("prioty = ?", 2)
	}})
	user.Scope(&admin.Scope{Name: beego.AppConfig.String("user"), Group: "权限等级", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Where("prioty = ?", 3)
	}})
}
