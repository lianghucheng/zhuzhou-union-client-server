package user

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor"
	"zhuzhou-union-client-server/utils"
)

func SetAdmin(adminConfig *admin.Admin) {
	user := adminConfig.AddResource(&models.User{})



	user.AddProcessor(&resource.Processor{
		Name: "process_user_data",
		Handler: func(val interface{}, values *resource.MetaValues, context *qor.Context) error {
			if user, ok := val.(*models.User); ok {
				user.Password = utils.Md5(user.Password) // do something...
			}
			return nil
		},
	})
}
