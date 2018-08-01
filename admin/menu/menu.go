package menu

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
)

func SetAdmin(admin *admin.Admin) {
	Aadmin.AddResource(&models.Menu{})
}
