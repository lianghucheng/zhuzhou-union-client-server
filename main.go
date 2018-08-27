package main

import (
	"github.com/astaxie/beego"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
	"github.com/qor/l10n"
	"os"
	"path/filepath"
	"zhuzhou-union-client-server/admin"
	"zhuzhou-union-client-server/models"
	"zhuzhou-union-client-server/pkg/LocalI18n"
	_ "zhuzhou-union-client-server/routers"
	_ "github.com/astaxie/beego/session/redis"
	"encoding/gob"
)

func init() {
	gob.Register(&models.User{})
	initI18n()
	initArgs()
	models.Connect()
}

func main() {
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
		if v == "-transfer" {
			models.Connect()
			//transfer()
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
	LocalI18n.LocalI18n = I18n
}
