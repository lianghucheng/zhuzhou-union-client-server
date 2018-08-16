package utils

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Weather struct {
	ShowapiResCode  int             `json:"showapi_res_code"`
	ShowapiResError string          `json:"showapi_res_error"`
	ShowapiResBody  *ShowapiResBody `json:"showapi_res_body"`
}

type ShowapiResBody struct {
	RetCode  int    `json:"ret_code"`
	Area     string `json:"area"`
	AreaId   string `json:"area_id"`
	HourList []*HourList
}

type HourList struct {
	WeatherCode   string `json:"weather_code"`
	Time          string `json:"time"`
	WindDirection string `json:"wind_direction"`
	WindPower     string `json:"wind_power"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
}

type ResultWeather struct{
	Update string
	Temperature string
	Weather string
	WindDirection string
	WindPower string
}

func TodayWeather() ResultWeather {
	var weather *Weather
	resultWeather:=ResultWeather{}
	req, err := http.NewRequest("GET", "http://saweather.market.alicloudapi.com/hour24?area=株洲", nil)
	req.Header.Add("Authorization", "APPCODE 0757f19aa0e84efb8822e6cd1bda9230")
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		beego.Debug(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Debug(err)
	}
	json.Unmarshal(body, &weather)
	nowTime := weather.ShowapiResBody.HourList[0].Time
	nowYear := substr(nowTime, 0, 4)
	nowMonth := substr(nowTime, 4, 6)
	nowDay := substr(nowTime, 6, 8)
	nowHour := substr(nowTime, 8, 10)
	nowMinute := substr(nowTime, 10, 12)
	resultWeather.Update=nowYear + "-" + nowMonth + "-" + nowDay + " " + nowHour + ":" + nowMinute
	resultWeather.Temperature=weather.ShowapiResBody.HourList[11].Temperature
	resultWeather.Weather=weather.ShowapiResBody.HourList[11].Weather
	resultWeather.WindDirection=weather.ShowapiResBody.HourList[11].WindDirection
	resultWeather.WindPower=weather.ShowapiResBody.HourList[11].WindPower
	return resultWeather
}

func substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}