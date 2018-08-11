package main

import (
	"github.com/astaxie/beego"
	"os"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/models"
	_ "zhuzhou-union-client-server/routers"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
	"path/filepath"
	"github.com/qor/l10n"
	"zhuzhou-union-client-server/pkg/LocalI18n"
	"zhuzhou-union-client-server/utils"
)

func init() {
	initI18n()
	initArgs()
	models.Connect()
}

func main() {
	utils.SendMsg("18374878791","123456")
	beego.Handler("/admin", admin.GetHandler(), true)
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.SyncDB()
			os.Exit(0)
		}
	}
}

func initI18n() {
	i18nPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	I18n := i18n.New(yaml.New(filepath.Join(i18nPath, "conf/en-US.yaml")))

	I18n.SaveTranslation(&i18n.Translation{Key: "qor_i18n.form.saved", Locale: "en-US", Value: "保存"})

	l10n.Global = "zh-CN"
	I18n.T("en-US", "demo.greeting") // Not exist at first
	I18n.T("en-US", "demo.hello")    // Exists in the yml file
	LocalI18n.LocalI18n = I18n
}
