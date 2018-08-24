package utils

import (
	"github.com/astaxie/beego"
	"regexp"
)

func MobileRegexp(mobile string) bool {
	matched, err := regexp.MatchString(beego.AppConfig.String("yidong"), mobile)
	if matched && err == nil {

		return true
	}
	matched, err = regexp.MatchString(beego.AppConfig.String("liantong"), mobile)
	if matched && err == nil {
		return true
	}
	matched, err = regexp.MatchString(beego.AppConfig.String("dianxin"), mobile)
	if matched && err == nil {
		return true
	}
	return false
}
