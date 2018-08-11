package utils

import (
	"github.com/lianghucheng/captcha"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

func GetCode(username string)string{
	id:=captcha.NewLen(6)
	byte_code:=captcha.GetCode(id)
	str_code:=""
	for _,v:=range byte_code{
		str_code+=strconv.Itoa(int(v))
	}
	beego.Debug(str_code)
	client:=redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	defer client.Close()
	client.HSet(username,username,id)
	return str_code
}

func VerifyCode(username,code string)bool{
	client:=redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	id,err:=client.HGet(username,username).Result()
	if err!=nil{
		beego.Error("不存在验证码key",err)
		return false
	}
	if !captcha.VerifyString(id,code){
		client.HDel(username,username)
		return false
	}
	return true
}