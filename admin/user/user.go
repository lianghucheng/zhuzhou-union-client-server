package user

import (
	"github.com/qor/admin"
	"zhuzhou-union-client-server/models"
)

func SetAdmin(admin *admin.Admin) {
	admin.AddResource(&models.User{})
}
