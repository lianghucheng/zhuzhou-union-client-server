package menu

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
	"github.com/qor/qor/resource"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"strings"
	"github.com/qor/validations"
)

func SetAdmin(adminConfig *admin.Admin) {
	menu := adminConfig.AddResource(&models.Menu{}, &admin.Config{PageCount: 10})

	menu.AddValidator(&resource.Validator{
		Name: "check_has_name",
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if meta := metaValues.Get("Name"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Name", "栏目名不能为空")
				}
			}
			return nil
		},
	})

}
