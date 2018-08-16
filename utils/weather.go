package utils

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

/*type Weather struct {
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

func GetWeather() ResultWeather {
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
}*/

type Weather struct {
	Date	string	`json:"date"`
	Message	string	`json:"message"`
	Status	int	`json:"status"`
	City	string	`json:"city"`
	Count	int	`json:"count"`
	Data	Data	`json:"data"`
}

type Data struct {
	Ganmao	string	`json:"ganmao"`
	Yesterday	Yesterday	`json:"yesterday"`
	Forecast	[]Forecast	`json:"forecast"`
	Shidu	string	`json:"shidu"`
	Pm25	float64	`json:"pm25"`
	Pm10	float64	`json:"pm10"`
	Quality	string	`json:"quality"`
	Wendu	string	`json:"wendu"`
}

type Yesterday struct {
	Sunset	string	`json:"sunset"`
	Fl	string	`json:"fl"`
	Notice	string	`json:"notice"`
	Date	string	`json:"date"`
	High	string	`json:"high"`
	Low	string	`json:"low"`
	Type	string	`json:"type"`
	Sunrise	string	`json:"sunrise"`
	Aqi	float64	`json:"aqi"`
	Fx	string	`json:"fx"`
}

type Forecast struct {
	Sunrise	string	`json:"sunrise"`
	Aqi	float64	`json:"aqi"`
	Fl	string	`json:"fl"`
	Type	string	`json:"type"`
	Date	string	`json:"date"`
	High	string	`json:"high"`
	Low	string	`json:"low"`
	Sunset	string	`json:"sunset"`
	Fx	string	`json:"fx"`
	Notice	string	`json:"notice"`
}

type ResultWeather struct{
	High string
	Low string
	Weather string
	WindDirection string
	WindPower string
	Notice string
}

func GetWeather()ResultWeather{
	resultWeather:=ResultWeather{}
	resp,err:=http.Get(`https://www.sojson.com/open/api/weather/json.shtml?city=株洲`)
	if err!=nil{
		beego.Error(err)
		return resultWeather
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
		return resultWeather

	}
	weather:=Weather{}
	if err:=json.Unmarshal(body,&weather);err!=nil{
		beego.Error(err)
		return resultWeather
	}
	resultWeather.High=weather.Data.Forecast[0].High
	resultWeather.Low=weather.Data.Forecast[0].Low
	resultWeather.Weather=weather.Data.Forecast[0].Type
	resultWeather.WindDirection=weather.Data.Forecast[0].Fx
	resultWeather.WindPower=weather.Data.Forecast[0].Fl
	resultWeather.Notice=weather.Data.Forecast[0].Notice
	beego.Debug(resultWeather)
	return resultWeather
}